package action

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/event"
	"github.com/pascalallen/carline/internal/carline/domain/password"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service/tokenservice"
	"github.com/pascalallen/carline/internal/carline/infrastructure/storage"
	"time"
)

type EmailsProjectionState struct {
	EmailAddresses map[string]string `json:"email_addresses"`
}

type LoginUserRules struct {
	EmailAddress string `form:"email_address" json:"email_address" binding:"required,max=100,email"`
	Password     string `form:"password" json:"password" binding:"required"`
}

type UserData struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
	CreatedAt    string `json:"created_at"`
	ModifiedAt   string `json:"modified_at"`
	DeletedAt    string `json:"deleted_at,omitempty"`
}

type ResponseData struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	User         UserData `json:"user"`
	Roles        []string `json:"roles,omitempty"`
	Permissions  []string `json:"permissions,omitempty"`
}

type RefreshTokensRules struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RegisterUserRules struct {
	FirstName       string `form:"first_name" json:"first_name" binding:"required,max=100"`
	LastName        string `form:"last_name" json:"last_name" binding:"required,max=100"`
	EmailAddress    string `form:"email_address" json:"email_address" binding:"required,max=100,email"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=Password"`
}

func HandleLoginUser(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request LoginUserRules

		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
			responder.BadRequestResponse(c, errors.New(errorMessage))

			return
		}

		getUserByEmailAddressQuery := query.GetUserByEmailAddress{EmailAddress: request.EmailAddress}
		result, err := queryBus.Fetch(getUserByEmailAddressQuery)
		if result == nil || err != nil {
			errorMessage := "invalid credentials"
			responder.UnauthorizedResponse(c, errors.New(errorMessage))

			return
		}

		u := result.(*user.User)

		if u.PasswordHash.Compare(request.Password) == false {
			errorMessage := "invalid credentials"
			responder.UnauthorizedResponse(c, errors.New(errorMessage))

			return
		}

		userClaims := tokenservice.UserClaims{
			Id:    u.Id.String(),
			First: u.FirstName,
			Last:  u.LastName,
			StandardClaims: jwt.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			},
		}

		signedAccessToken, err := tokenservice.NewAccessToken(userClaims)
		if err != nil {
			errorMessage := "error creating access token"
			responder.InternalServerErrorResponse(c, errors.New(errorMessage))

			return
		}

		var roles []string
		for _, r := range u.Roles {
			roles = append(roles, r.Name)
		}

		var permissions []string
		for _, p := range u.Permissions() {
			permissions = append(permissions, p.Name)
		}

		userData := UserData{
			Id:           u.Id.String(),
			FirstName:    u.FirstName,
			LastName:     u.LastName,
			EmailAddress: u.EmailAddress,
			CreatedAt:    u.CreatedAt.String(),
			ModifiedAt:   u.ModifiedAt.String(),
		}
		if !u.DeletedAt.IsZero() {
			userData.DeletedAt = u.DeletedAt.String()
		}

		refreshClaims := jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		}

		signedRefreshToken, err := tokenservice.NewRefreshToken(refreshClaims)
		if err != nil {
			errorMessage := "error creating refresh token"
			responder.InternalServerErrorResponse(c, errors.New(errorMessage))

			return
		}

		responseData := &ResponseData{
			AccessToken:  signedAccessToken,
			RefreshToken: signedRefreshToken,
			User:         userData,
			Roles:        roles,
			Permissions:  permissions,
		}

		responder.CreatedResponse(c, responseData)

		return
	}
}

func HandleRefreshTokens(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request RefreshTokensRules

		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
			responder.BadRequestResponse(c, errors.New(errorMessage))

			return
		}

		userClaims := tokenservice.ParseAccessToken(request.AccessToken)
		refreshClaims := tokenservice.ParseRefreshToken(request.RefreshToken)

		getUserById := query.GetUserById{Id: ulid.MustParse(userClaims.Id)}
		result, err := queryBus.Fetch(getUserById)
		if result == nil || err != nil {
			errorMessage := "invalid credentials"
			responder.UnauthorizedResponse(c, errors.New(errorMessage))

			return
		}

		u := result.(*user.User)

		// refresh token is expired
		if refreshClaims.Valid() != nil {
			request.RefreshToken, err = tokenservice.NewRefreshToken(*refreshClaims)
			if err != nil {
				errorMessage := "error creating refresh token"
				responder.InternalServerErrorResponse(c, errors.New(errorMessage))

				return
			}
		}

		// access token is expired
		if userClaims.StandardClaims.Valid() != nil && refreshClaims.Valid() == nil {
			request.AccessToken, err = tokenservice.NewAccessToken(*userClaims)
			if err != nil {
				errorMessage := "error creating access token"
				responder.InternalServerErrorResponse(c, errors.New(errorMessage))

				return
			}
		}

		var roles []string
		for _, r := range u.Roles {
			roles = append(roles, r.Name)
		}

		var permissions []string
		for _, p := range u.Permissions() {
			permissions = append(permissions, p.Name)
		}

		userData := UserData{
			Id:           ulid.ULID(u.Id).String(),
			FirstName:    u.FirstName,
			LastName:     u.LastName,
			EmailAddress: u.EmailAddress,
			CreatedAt:    u.CreatedAt.String(),
			ModifiedAt:   u.ModifiedAt.String(),
		}
		if !u.DeletedAt.IsZero() {
			userData.DeletedAt = u.DeletedAt.String()
		}

		responseData := &ResponseData{
			AccessToken:  request.AccessToken,
			RefreshToken: request.RefreshToken,
			User:         userData,
			Roles:        roles,
			Permissions:  permissions,
		}

		responder.CreatedResponse(c, responseData)

		return
	}
}

func HandleRegisterUser(eventStore storage.EventStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request RegisterUserRules

		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
			responder.BadRequestResponse(c, errors.New(errorMessage))

			return
		}

		var result EmailsProjectionState
		if err := eventStore.UnmarshalProjectionResult("user-email-addresses", &result); err != nil {
			errorMessage := fmt.Sprint("Something went wrong. If you already have an account, please log in.")
			responder.UnprocessableEntityResponse(c, errors.New(errorMessage))

			return
		}

		for emailAddress := range result.EmailAddresses {
			if emailAddress == request.EmailAddress {
				errorMessage := fmt.Sprint("Something went wrong. If you already have an account, please log in.")
				responder.UnprocessableEntityResponse(c, errors.New(errorMessage))

				return
			}
		}

		registerEvent := event.UserRegistered{
			Id:           ulid.Make(),
			FirstName:    request.FirstName,
			LastName:     request.LastName,
			EmailAddress: request.EmailAddress,
			PasswordHash: password.Create(request.Password),
		}
		streamId := fmt.Sprintf("user-%s", registerEvent.Id)
		err := eventStore.AppendToStream(streamId, registerEvent)
		if err != nil {
			errorMessage := fmt.Sprint("Something went wrong. If you already have an account, please log in.")
			responder.UnprocessableEntityResponse(c, errors.New(errorMessage))

			return
		}

		responder.CreatedResponse[RegisterUserRules](c, &request)

		return
	}
}
