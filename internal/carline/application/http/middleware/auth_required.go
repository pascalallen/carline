package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service"
	"strings"
)

func AuthRequired(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			responder.BadRequestResponse(c, errors.New("invalid authorization header"))
			return
		}

		accessToken := parts[1]
		userClaims := service.ParseAccessToken(accessToken)
		if userClaims == nil {
			responder.UnauthorizedResponse(c, errors.New("invalid or expired token"))
			return
		}

		q := query.GetUserById{Id: ulid.MustParse(userClaims.Id)}
		result, err := queryBus.Fetch(q)
		u, ok := result.(*user.User)
		if u == nil || err != nil || !ok {
			responder.UnauthorizedResponse(c, errors.New("invalid credentials"))
			return
		}

		c.Set("userId", u.Id)

		c.Next()
	}
}
