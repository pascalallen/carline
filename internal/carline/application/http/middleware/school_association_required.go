package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service/tokenservice"
	"strings"
)

func SchoolAssociationRequired(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		schoolId := c.Param("schoolId")

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responder.UnauthorizedResponse(c, errors.New("authorization header required"))
			return
		}

		accessToken := strings.Split(authHeader, " ")[1]
		userClaims := tokenservice.ParseAccessToken(accessToken)
		userId, err := ulid.Parse(userClaims.Id)
		if err != nil {
			responder.UnauthorizedResponse(c, errors.New("invalid user ID in token"))
			return
		}

		q := query.ListSchools{UserId: userId}
		result, err := queryBus.Fetch(q)
		if err != nil {
			responder.InternalServerErrorResponse(c, errors.New("something went wrong when fetching user schools"))
			return
		}

		schools, ok := result.(*[]school.School)
		if !ok {
			responder.InternalServerErrorResponse(c, errors.New("failed to parse the list of schools"))
			return
		}

		isAssociated := false
		for _, s := range *schools {
			if s.Id.String() == schoolId {
				isAssociated = true
				break
			}
		}

		if !isAssociated {
			responder.ForbiddenResponse(c, errors.New("user not associated with school"))
			return
		}

		c.Next()
	}
}
