package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service/tokenservice"
	"strings"
)

func AuthRequired(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responder.BadRequestResponse(c, errors.New("authorization header is required"))

			return
		}

		accessToken := strings.Split(authHeader, " ")[1]
		userClaims := tokenservice.ParseAccessToken(accessToken)

		getUserById := query.GetUserById{Id: ulid.MustParse(userClaims.Id)}
		result, err := queryBus.Fetch(getUserById)
		if result == nil || err != nil {
			errorMessage := "invalid credentials"
			responder.UnauthorizedResponse(c, errors.New(errorMessage))

			return
		}

		c.Next()
	}
}
