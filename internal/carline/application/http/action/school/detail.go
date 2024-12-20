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

type DetailResponsePayload struct {
	School school.School `json:"school"`
}

func HandleDetail(queryBus messaging.QueryBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		schoolId := c.Param("schoolId")

		authHeader := c.GetHeader("Authorization")
		accessToken := strings.Split(authHeader, " ")[1]
		userClaims := service.ParseAccessToken(accessToken)
		userId := ulid.MustParse(userClaims.Id)

		q := query.GetSchoolByIdAndUserId{UserId: userId, Id: ulid.MustParse(schoolId)}
		result, err := queryBus.Fetch(q)
		s, ok := result.(*school.School)
		if err != nil || !ok {
			errorMessage := fmt.Sprint("Something went wrong.")
			responder.UnprocessableEntityResponse(c, errors.New(errorMessage))

			return
		}

		response := DetailResponsePayload{
			School: *s,
		}
		responder.OkResponse[DetailResponsePayload](c, &response)

		return

	}
}
