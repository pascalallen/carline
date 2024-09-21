package routes

import (
	school2 "github.com/pascalallen/carline/internal/carline/application/http/action/school"
	"github.com/pascalallen/carline/internal/carline/application/http/middleware"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

func (r Router) Schools(userRepo user.Repository, schoolRepo school.Repository, commandBus messaging.CommandBus) {
	v := r.engine.Group(v1)
	{
		v.POST(
			"/schools",
			middleware.AuthRequired(userRepo),
			school2.HandleCreate(schoolRepo, commandBus),
		)

		v.GET(
			"/schools",
			middleware.AuthRequired(userRepo),
			school2.HandleList(schoolRepo),
		)
	}
}
