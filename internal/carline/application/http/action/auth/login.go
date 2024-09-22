package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service/tokenservice"
	"time"
)

type LoginRequestPayload struct {
	EmailAddress string `form:"email_address" json:"email_address" binding:"required,max=100,email"`
	Password     string `form:"password" json:"password" binding:"required"`
}

func HandleLoginUser(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request LoginRequestPayload

		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
			responder.BadRequestResponse(c, errors.New(errorMessage))

			return
		}

		q := query.GetUserByEmailAddress{EmailAddress: request.EmailAddress}
		result, err := queryBus.Fetch(q)
		u, ok := result.(*user.User)
		if u == nil || err != nil || !ok {
			errorMessage := "invalid credentials"
			responder.UnauthorizedResponse(c, errors.New(errorMessage))

			return
		}

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
		}

		if u.ModifiedAt != nil {
			userData.ModifiedAt = u.ModifiedAt.String()
		}

		if u.DeletedAt != nil {
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

		responseData := &TokenResponsePayload{
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
