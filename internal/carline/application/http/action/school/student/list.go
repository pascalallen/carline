package student

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/student"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type ListRequestPayload struct {
	IncludeDeleted bool `form:"include_deleted" json:"include_deleted"`
}

type ListResponsePayload struct {
	Students []student.Student `json:"students"`
}

func HandleList(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("schoolId")
		if id == "" {
			responder.BadRequestResponse(c, errors.New("school ID required"))

			return
		}

		var request ListRequestPayload
		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
			responder.BadRequestResponse(c, errors.New(errorMessage))

			return
		}

		q := query.ListStudents{
			SchoolId:       ulid.MustParse(id),
			IncludeDeleted: request.IncludeDeleted,
		}
		result, err := queryBus.Fetch(q)
		s, ok := result.(*[]student.Student)
		if err != nil || !ok {
			errorMessage := fmt.Sprint("Something went wrong.")
			responder.UnprocessableEntityResponse(c, errors.New(errorMessage))

			return
		}

		response := ListResponsePayload{
			Students: *s,
		}
		responder.OkResponse[ListResponsePayload](c, &response)

		return

	}
}
