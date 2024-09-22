package routes

import (
	school2 "github.com/pascalallen/carline/internal/carline/application/http/action/school"
	"github.com/pascalallen/carline/internal/carline/application/http/middleware"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

func (r Router) Schools(queryBus messaging.QueryBus, commandBus messaging.CommandBus) {
	v := r.engine.Group(v1)
	{
		v.POST(
			"/schools",
			middleware.AuthRequired(queryBus),
			school2.HandleCreate(queryBus, commandBus),
		)

		v.GET(
			"/schools",
			middleware.AuthRequired(queryBus),
			school2.HandleList(queryBus),
		)
	}
}
