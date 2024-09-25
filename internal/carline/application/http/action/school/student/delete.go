package student

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/student"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type DeleteResponsePayload struct {
	Id string `json:"id"`
}

func HandleDelete(queryBus messaging.QueryBus, commandBus messaging.CommandBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		schoolId := c.Param("schoolId")
		studentId := c.Param("studentId")

		if schoolId == "" {
			responder.BadRequestResponse(c, errors.New("school ID required"))

			return
		}

		if studentId == "" {
			responder.BadRequestResponse(c, errors.New("student ID required"))

			return
		}

		q := query.GetStudentById{SchoolId: ulid.MustParse(schoolId), Id: ulid.MustParse(studentId)}
		result, err := queryBus.Fetch(q)
		s, ok := result.(*student.Student)
		if s == nil || err != nil || !ok {
			responder.NotFoundResponse(c, errors.New("student not found"))

			return
		}

		cmd := command.DeleteStudent{SchoolId: q.SchoolId, Id: q.Id}
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
