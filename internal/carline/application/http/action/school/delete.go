package school

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type DeleteResponsePayload struct {
	Id string `json:"id"`
}

func HandleDelete(queryBus messaging.QueryBus, commandBus messaging.CommandBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if id == "" {
			responder.BadRequestResponse(c, errors.New("school ID required"))

			return
		}

		q := query.GetSchoolById{Id: ulid.MustParse(id)}
		result, err := queryBus.Fetch(q)
		s, ok := result.(*school.School)
		if s == nil || err != nil || !ok {
			responder.NotFoundResponse(c, errors.New("school not found"))

			return
		}

		cmd := command.RemoveSchool{Id: q.Id}
		err = commandBus.Execute(cmd)
		if err != nil {
			responder.InternalServerErrorResponse(c, err)

			return
		}

		response := DeleteResponsePayload{
			Id: s.Id.String(),
		}
		responder.OkResponse[DeleteResponsePayload](c, &response)

		return
	}
}
