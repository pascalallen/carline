package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type CreateRequestPayload struct {
	FirstName    string `form:"first_name" json:"first_name" binding:"required,max=100"`
	LastName     string `form:"last_name" json:"last_name" binding:"required,max=100"`
	EmailAddress string `form:"email_address" json:"email_address" binding:"required,max=100"`
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

		q := query.GetUserByEmailAddress{EmailAddress: request.EmailAddress}
		result, err := queryBus.Fetch(q)
		s, ok := result.(*user.User)
		if s != nil || err != nil || !ok {
			errorMessage := fmt.Sprint("User already exists.")
			responder.UnprocessableEntityResponse(c, errors.New(errorMessage))

			return
		}

		schoolId := ulid.MustParse(c.Param("schoolId"))
		cmd := command.RegisterUser{
			Id:           ulid.Make(),
			SchoolId:     &schoolId,
			FirstName:    request.FirstName,
			LastName:     request.LastName,
			EmailAddress: request.EmailAddress,
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
