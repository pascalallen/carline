package school

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service/tokenservice"
	"strings"
)

type CreateRequestPayload struct {
	Name string `form:"name" json:"name" binding:"required,max=100"`
}

type CreatedResponsePayload struct {
	Id ulid.ULID `json:"id"`
}

func HandleCreate(queryBus messaging.QueryBus, commandBus messaging.CommandBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request CreateRequestPayload

		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
			responder.BadRequestResponse(c, errors.New(errorMessage))

			return
		}

		q := query.GetSchoolByName{Name: request.Name}
		result, err := queryBus.Fetch(q)
		s, ok := result.(*school.School)
		if s != nil || err != nil || !ok {
			errorMessage := fmt.Sprint("School already exists.")
			responder.UnprocessableEntityResponse(c, errors.New(errorMessage))

			return
		}

		authHeader := c.GetHeader("Authorization")
		accessToken := strings.Split(authHeader, " ")[1]
		userClaims := tokenservice.ParseAccessToken(accessToken)
		userId := ulid.MustParse(userClaims.Id)

		cmd := command.CreateSchool{
			Id:     ulid.Make(),
			Name:   request.Name,
			UserId: userId,
		}
		err = commandBus.Execute(cmd)
		if err != nil {
			errorMessage := fmt.Sprintf("Something went wrong executing the command: %s", err.Error())
			responder.InternalServerErrorResponse(c, errors.New(errorMessage))
			return
		}

		response := CreatedResponsePayload{
			Id: cmd.Id,
		}
		responder.CreatedResponse[CreatedResponsePayload](c, &response)

		return
	}
}
