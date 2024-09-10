package routes

import (
	"github.com/pascalallen/carline/internal/carline/application/http/action"
	"github.com/pascalallen/carline/internal/carline/application/http/middleware"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

func (r Router) Temp(queryBus messaging.QueryBus) {
	v := r.engine.Group(v1)
	{
		v.GET(
			"/temp",
			middleware.AuthRequired(queryBus),
			action.HandleTemp(),
		)

		ch := make(chan string)
		v.POST(
			"/event-stream",
			middleware.AuthRequired(queryBus),
			action.HandleEventStreamPost(ch),
		)
		v.GET(
			"/event-stream",
			middleware.EventStreamHeaders(),
			action.HandleEventStreamGet(ch),
		)
	}
}
