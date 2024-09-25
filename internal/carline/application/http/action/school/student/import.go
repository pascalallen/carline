package student

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/http/responder"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"io"
)

func HandleImport(commandBus messaging.CommandBus) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			responder.BadRequestResponse(c, errors.New("school ID required"))

			return
		}

		if err := c.Request.ParseMultipartForm(2 << 20); err != nil {
			responder.BadRequestResponse(c, errors.New("file size must be less than 2MB"))

			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			responder.BadRequestResponse(c, errors.New("file required"))

			return
		}

		srcFile, err := file.Open()
		if err != nil {
			responder.InternalServerErrorResponse(c, errors.New("failed to open uploaded file"))
			return
		}
		defer srcFile.Close()

		fileBuffer, err := io.ReadAll(srcFile)
		if err != nil {
			responder.InternalServerErrorResponse(c, errors.New("failed to read uploaded file"))
			return
		}

		cmd := command.ImportStudents{
			SchoolId:   ulid.MustParse(id),
			FileBuffer: fileBuffer,
		}
		err = commandBus.Execute(cmd)
		if err != nil {
			errorMessage := fmt.Sprintf("Something went wrong executing the command: %s", err.Error())
			responder.InternalServerErrorResponse(c, errors.New(errorMessage))

			return
		}

		responder.CreatedResponse[string](c, &id)

		return
	}
}
