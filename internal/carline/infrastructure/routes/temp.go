package routes

import (
	"github.com/pascalallen/carline/internal/carline/application/http/action"
	"github.com/pascalallen/carline/internal/carline/application/http/middleware"
	"github.com/pascalallen/carline/internal/carline/domain/user"
)

func (r Router) Temp(repository user.UserRepository) {
	v := r.engine.Group(v1)
	{
		v.GET(
			"/temp",
			middleware.AuthRequired(repository),
			action.HandleTemp(),
		)

		ch := make(chan string)
		v.POST(
			"/event-stream",
			middleware.AuthRequired(repository),
			action.HandleEventStreamPost(ch),
		)
		v.GET(
			"/event-stream",
			middleware.EventStreamHeaders(),
			action.HandleEventStreamGet(ch),
		)
	}
}
