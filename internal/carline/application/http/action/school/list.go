package school

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service"
	"strings"
)

type ListResponsePayload struct {
	Schools []school.School `json:"schools"`
}

func HandleList(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		accessToken := strings.Split(authHeader, " ")[1]
		userClaims := service.ParseAccessToken(accessToken)
		userId := ulid.MustParse(userClaims.Id)

		q := query.ListSchools{UserId: userId}
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
