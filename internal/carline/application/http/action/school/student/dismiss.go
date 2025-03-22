package student

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/websocket"
)

type DismissalRequestPayload struct {
	TagNumber string `json:"tag_number" binding:"required,max=100"`
}

type DismissalResponsePayload struct {
	TagNumber string `json:"tag_number"`
}

func HandleDismissal(commandBus messaging.CommandBus, websocketHub *websocket.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request DismissalRequestPayload

		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
			responder.BadRequestResponse(c, errors.New(errorMessage))

			return
		}

		schoolId := c.Param("schoolId")

		cmd := command.DismissStudents{
			SchoolId:  ulid.MustParse(schoolId),
			TagNumber: request.TagNumber,
		}
		err := commandBus.Execute(cmd)
		if err != nil {
			errorMessage := fmt.Sprintf("Something went wrong executing the command: %s", err.Error())
			responder.InternalServerErrorResponse(c, errors.New(errorMessage))
			return
		}

		websocket.ServeWs(websocketHub, ulid.MustParse(schoolId), c)

		response := DismissalResponsePayload{
			TagNumber: request.TagNumber,
		}
		responder.CreatedResponse[DismissalResponsePayload](c, &response)

		return

	}
}
