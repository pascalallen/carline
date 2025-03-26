package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service"
	"log"
	"strings"
)

func SchoolAssociationRequired(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bypass WebSocket requests:
		if c.Request.Header.Get("Upgrade") == "websocket" {
			log.Println("Bypassing SchoolAssociationRequired for WebSocket request")
			c.Next()
			return
		}

		schoolId := c.Param("schoolId")

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responder.UnauthorizedResponse(c, errors.New("authorization header required"))
			return
		}

		accessToken := strings.Split(authHeader, " ")[1]
		userClaims := service.ParseAccessToken(accessToken)
		userId, err := ulid.Parse(userClaims.Id)
		if err != nil {
			responder.UnauthorizedResponse(c, errors.New("invalid user ID in token"))
			return
		}

		q := query.GetSchoolByIdAndUserId{UserId: userId, Id: ulid.MustParse(schoolId)}
		result, err := queryBus.Fetch(q)
		if err != nil {
			responder.InternalServerErrorResponse(c, errors.New("something went wrong fetching school for user"))
			return
		}

		associatedSchool, ok := result.(*school.School)
		if !ok {
			responder.InternalServerErrorResponse(c, errors.New("failed to parse school"))
			return
		}

		if associatedSchool == nil {
			responder.ForbiddenResponse(c, errors.New("user not associated with school"))
			return
		}

		c.Next()
	}
}
