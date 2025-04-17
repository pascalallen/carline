package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

func SchoolAssociationRequired(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdRaw, exists := c.Get("userId")
		if !exists {
			responder.UnauthorizedResponse(c, errors.New("user not authenticated"))
			return
		}

		userId, ok := userIdRaw.(ulid.ULID)
		if !ok {
			responder.InternalServerErrorResponse(c, errors.New("failed to retrieve user ID"))
			return
		}

		schoolId := c.Param("schoolId")
		schoolULID, err := ulid.Parse(schoolId)
		if err != nil {
			responder.BadRequestResponse(c, errors.New("invalid school ID"))
			return
		}

		q := query.GetSchoolByIdAndUserId{UserId: userId, Id: schoolULID}
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
