package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service/tokenservice"
	"strings"
)

func AuthRequired(userRepository user.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responder.BadRequestResponse(c, errors.New("authorization header is required"))

			return
		}

		accessToken := strings.Split(authHeader, " ")[1]
		userClaims := tokenservice.ParseAccessToken(accessToken)

		u, err := userRepository.GetById(ulid.MustParse(userClaims.Id))
		if u == nil || err != nil {
			responder.UnauthorizedResponse(c, errors.New("invalid credentials"))

			return
		}

		c.Next()
	}
}
