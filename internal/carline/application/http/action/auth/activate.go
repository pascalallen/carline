package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/crypto"
	"github.com/pascalallen/carline/internal/carline/domain/password"
	"github.com/pascalallen/carline/internal/carline/domain/security_token"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service"
	"time"
)

type ActivateRequestPayload struct {
	Token           string `form:"token" json:"token" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=Password"`
}

func HandleActivateUser(queryBus messaging.QueryBus, commandBus messaging.CommandBus, securityTokenService security_token.Service, userRepo user.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request ActivateRequestPayload

		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
			responder.BadRequestResponse(c, errors.New(errorMessage))
			return
		}

		securityToken, err := securityTokenService.FetchToken(crypto.Crypto(request.Token))
		if err != nil {
			responder.BadRequestResponse(c, errors.New("invalid or expired token"))
			return
		}

		if securityToken.IsExpired() {
			responder.BadRequestResponse(c, errors.New("security token has expired"))
			return
		}

		q := query.GetUserById{Id: securityToken.UserId}
		result, err := queryBus.Fetch(q)
		u, ok := result.(*user.User)
		if err != nil || !ok || u == nil {
			responder.BadRequestResponse(c, errors.New("user not found"))
			return
		}

		if request.Password != request.ConfirmPassword {
			responder.BadRequestResponse(c, errors.New("password and confirmation do not match"))
			return
		}

		passwordHash := password.Create(request.Password)
		u.SetPasswordHash(passwordHash)
		err = userRepo.Save(u)
		if err != nil {
			responder.InternalServerErrorResponse(c, errors.New("failed to update user password"))
			return
		}

		err = securityTokenService.ClearTokensForUser(*u)
		if err != nil {
			responder.InternalServerErrorResponse(c, errors.New("failed to clear security tokens"))
			return
		}

		userClaims := service.UserClaims{
			Id:    u.Id.String(),
			First: u.FirstName,
			Last:  u.LastName,
			StandardClaims: jwt.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			},
		}

		signedAccessToken, err := service.NewAccessToken(userClaims)
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

		refreshClaims := jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		}

		signedRefreshToken, err := service.NewRefreshToken(refreshClaims)
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
