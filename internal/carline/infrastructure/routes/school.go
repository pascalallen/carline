package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/application/http/action/school"
	"github.com/pascalallen/carline/internal/carline/application/http/action/school/student"
	"github.com/pascalallen/carline/internal/carline/application/http/action/school/user"
	"github.com/pascalallen/carline/internal/carline/application/http/middleware"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/websocket"
)

func (r Router) Schools(queryBus messaging.QueryBus, commandBus messaging.CommandBus, websocketHub *websocket.Hub) {
	v := r.engine.Group(v1)
	{
		v.POST(
			"/schools",
			middleware.AuthRequired(queryBus),
			school.HandleCreate(queryBus, commandBus),
		)

		v.GET(
			"/schools",
			middleware.AuthRequired(queryBus),
			school.HandleList(queryBus),
		)

		v.GET(
			"/schools/:schoolId",
			middleware.AuthRequired(queryBus),
			middleware.SchoolAssociationRequired(queryBus),
			school.HandleDetail(queryBus),
		)

		v.DELETE(
			"/schools/:schoolId",
			middleware.AuthRequired(queryBus),
			middleware.SchoolAssociationRequired(queryBus),
			school.HandleDelete(commandBus),
		)

		v.GET(
			"/schools/:schoolId/students",
			middleware.AuthRequired(queryBus),
			middleware.SchoolAssociationRequired(queryBus),
			student.HandleList(queryBus),
		)

		v.POST(
			"/schools/:schoolId/students/dismissals",
			middleware.AuthRequired(queryBus),
			middleware.SchoolAssociationRequired(queryBus),
			student.HandleDismissal(commandBus, websocketHub),
		)

		v.GET(
			"/schools/:schoolId/students/dismissals/ws",
			func(c *gin.Context) {
				//middleware.AuthRequired(queryBus) // TODO
				//middleware.SchoolAssociationRequired(queryBus) // TODO
				schoolId := c.Param("schoolId")
				groupId := ulid.MustParse(schoolId)
				websocket.ServeWs(websocketHub, groupId, c)
			},
		)

		v.POST(
			"/schools/:schoolId/students/import",
			middleware.AuthRequired(queryBus),
			middleware.SchoolAssociationRequired(queryBus),
			student.HandleImport(commandBus),
		)

		v.DELETE(
			"/schools/:schoolId/students/:studentId",
			middleware.AuthRequired(queryBus),
			middleware.SchoolAssociationRequired(queryBus),
			student.HandleDelete(queryBus, commandBus),
		)

		v.POST(
			"/schools/:schoolId/users",
			middleware.AuthRequired(queryBus),
			middleware.SchoolAssociationRequired(queryBus),
			user.HandleCreate(queryBus, commandBus),
		)

		v.GET(
			"/schools/:schoolId/users",
			middleware.AuthRequired(queryBus),
			middleware.SchoolAssociationRequired(queryBus),
			user.HandleList(queryBus),
		)

		v.GET(
			"/schools/:schoolId/users/:userId",
			middleware.AuthRequired(queryBus),
			middleware.SchoolAssociationRequired(queryBus),
			// TODO user.HandleDetail(queryBus),
		)

		v.PUT(
			"/schools/:schoolId/users/:userId",
			middleware.AuthRequired(queryBus),
			middleware.SchoolAssociationRequired(queryBus),
			// TODO user.HandleUpdate(commandBus),
		)

		v.DELETE(
			"/schools/:schoolId/users/:userId",
			middleware.AuthRequired(queryBus),
			middleware.SchoolAssociationRequired(queryBus),
			// TODO user.HandleDelete(commandBus),
		)
	}
}
