package school

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type ListRequestPayload struct {
	IncludeDeleted bool `form:"include_deleted" json:"include_deleted"`
}

type ListResponsePayload struct {
	Schools []school.School `json:"schools"`
}

func HandleList(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request ListRequestPayload

		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
			responder.BadRequestResponse(c, errors.New(errorMessage))

			return
		}

		q := query.ListSchools{IncludeDeleted: request.IncludeDeleted}
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
