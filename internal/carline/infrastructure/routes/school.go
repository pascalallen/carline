package routes

import (
	"github.com/pascalallen/carline/internal/carline/application/http/action/school"
	"github.com/pascalallen/carline/internal/carline/application/http/action/school/student"
	"github.com/pascalallen/carline/internal/carline/application/http/middleware"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

func (r Router) Schools(queryBus messaging.QueryBus, commandBus messaging.CommandBus) {
	v := r.engine.Group(v1)
	{
		v.POST(
			"/schools",
			middleware.AuthRequired(queryBus),
			school.HandleCreate(queryBus, commandBus),
		)

		v.DELETE(
			"/schools/:schoolId",
			middleware.AuthRequired(queryBus),
			school.HandleDelete(queryBus, commandBus),
		)

		v.GET(
			"/schools",
			middleware.AuthRequired(queryBus),
			school.HandleList(queryBus),
		)

		v.POST(
			"/schools/:schoolId/students/import",
			middleware.AuthRequired(queryBus),
			student.HandleImport(commandBus),
		)

		v.DELETE(
			"/schools/:schoolId/students/:studentId",
			middleware.AuthRequired(queryBus),
			student.HandleDelete(queryBus, commandBus),
		)

		v.GET(
			"/schools/:schoolId/students",
			middleware.AuthRequired(queryBus),
			student.HandleList(queryBus),
		)
	}
}
