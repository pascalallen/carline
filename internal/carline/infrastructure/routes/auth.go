package routes

import (
	"github.com/pascalallen/carline/internal/carline/application/http/action/auth"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

func (r Router) Auth(queryBus messaging.QueryBus, commandBus messaging.CommandBus) {
	v := r.engine.Group(v1)
	{
		a := v.Group("/auth")
		{
			a.POST("/register", auth.HandleRegisterUser(queryBus, commandBus))
			a.POST("/login", auth.HandleLoginUser(queryBus))
			a.PATCH("/refresh", auth.HandleRefreshTokens(queryBus))
			// router.POST("/request-reset", auth.HandleRequestPasswordReset)
			// router.POST("/reset-password", auth.HandleResetPassword)
		}
	}
}
