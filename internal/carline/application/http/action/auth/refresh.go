package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service"
)

type RefreshRequestPayload struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func HandleRefreshTokens(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request RefreshRequestPayload

		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
			responder.BadRequestResponse(c, errors.New(errorMessage))

			return
		}

		userClaims := service.ParseAccessToken(request.AccessToken)
		refreshClaims := service.ParseRefreshToken(request.RefreshToken)

		q := query.GetUserById{Id: ulid.MustParse(userClaims.Id)}
		result, err := queryBus.Fetch(q)
		u, ok := result.(*user.User)
		if u == nil || err != nil || !ok {
			errorMessage := "invalid credentials"
			responder.UnauthorizedResponse(c, errors.New(errorMessage))

			return
		}

		// refresh token is expired
		if refreshClaims.Valid() != nil {
			request.RefreshToken, err = service.NewRefreshToken(*refreshClaims)
			if err != nil {
				errorMessage := "error creating refresh token"
				responder.InternalServerErrorResponse(c, errors.New(errorMessage))

				return
			}
		}

		// access token is expired
		if userClaims.StandardClaims.Valid() != nil && refreshClaims.Valid() == nil {
			request.AccessToken, err = service.NewAccessToken(*userClaims)
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
			Id:           u.Id.String(),
			FirstName:    u.FirstName,
			LastName:     u.LastName,
			EmailAddress: u.EmailAddress,
			CreatedAt:    u.CreatedAt.String(),
		}

		if u.ModifiedAt != nil {
			userData.ModifiedAt = u.ModifiedAt.String()
		}

		responseData := &TokenResponsePayload{
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
