package routes

import (
	"github.com/pascalallen/carline/internal/carline/application/http/action"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/storage"
)

func (r Router) Auth(queryBus messaging.QueryBus, eventStore storage.EventStore) {
	v := r.engine.Group(v1)
	{
		auth := v.Group("/auth")
		{
			auth.POST("/register", action.HandleRegisterUser(eventStore))
			auth.POST("/login", action.HandleLoginUser(queryBus))
			auth.PATCH("/refresh", action.HandleRefreshTokens(queryBus))
			// router.POST("/request-reset", auth.HandleRequestPasswordReset)
			// router.POST("/reset-password", auth.HandleResetPassword)
		}
	}
}
