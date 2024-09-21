package school

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	school2 "github.com/pascalallen/carline/internal/carline/domain/school"
)

type ListRequestPayload struct {
	IncludeDeleted bool `form:"include_deleted" json:"include_deleted"`
}

type ListResponsePayload struct {
	Schools []school2.School `json:"schools"`
}

func HandleList(schoolRepository school2.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request ListRequestPayload

		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("Request validation error: %s", err.Error())
			responder.BadRequestResponse(c, errors.New(errorMessage))

			return
		}

		s, err := schoolRepository.GetAll(request.IncludeDeleted)
		if s != nil || err != nil {
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
