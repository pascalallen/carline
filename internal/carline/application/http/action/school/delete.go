package school

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service"
	"strings"
)

type DeleteResponsePayload struct {
	Id string `json:"id"`
}

func HandleDelete(commandBus messaging.CommandBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("schoolId")

		if id == "" {
			responder.BadRequestResponse(c, errors.New("school ID required"))

			return
		}

		authHeader := c.GetHeader("Authorization")
		accessToken := strings.Split(authHeader, " ")[1]
		userClaims := service.ParseAccessToken(accessToken)
		userId := ulid.MustParse(userClaims.Id)

		cmd := command.DeleteSchool{
			Id:     ulid.MustParse(id),
			UserId: userId,
		}
		err := commandBus.Execute(cmd)
		if err != nil {
			responder.InternalServerErrorResponse(c, err)

			return
		}

		response := DeleteResponsePayload{
			Id: id,
		}
		responder.OkResponse[DeleteResponsePayload](c, &response)

		return
	}
}
