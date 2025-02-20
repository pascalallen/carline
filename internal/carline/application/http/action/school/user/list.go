package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type ListResponsePayload struct {
	Users []user.User `json:"users"`
}

func HandleList(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("schoolId")
		if id == "" {
			responder.BadRequestResponse(c, errors.New("school ID required"))

			return
		}

		q := query.ListUsers{
			SchoolId: ulid.MustParse(id),
		}
		result, err := queryBus.Fetch(q)
		u, ok := result.(*[]user.User)
		if err != nil || !ok {
			errorMessage := fmt.Sprint("Something went wrong.")
			responder.UnprocessableEntityResponse(c, errors.New(errorMessage))

			return
		}

		response := ListResponsePayload{
			Users: *u,
		}
		responder.OkResponse[ListResponsePayload](c, &response)

		return

	}
}
