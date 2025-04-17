package school

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type ListResponsePayload struct {
	Schools []school.School `json:"schools"`
}

func HandleList(queryBus messaging.QueryBus) gin.HandlerFunc {
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

		q := query.ListSchools{UserId: userId}
		result, err := queryBus.Fetch(q)
		s, ok := result.(*[]school.School)
		if err != nil || !ok {
			errorMessage := fmt.Sprint("Something went wrong.")
			responder.UnprocessableEntityResponse(c, errors.New(errorMessage))
			return
		}

		response := ListResponsePayload{
			Schools: *s,
		}
		responder.OkResponse[ListResponsePayload](c, &response)

		return
	}
}
