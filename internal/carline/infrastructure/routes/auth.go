package routes

import (
	"github.com/pascalallen/carline/internal/carline/application/http/action"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

func (r Router) Auth(repository user.UserRepository, bus messaging.CommandBus) {
	v := r.engine.Group(v1)
	{
		auth := v.Group("/auth")
		{
			auth.POST("/register", action.HandleRegisterUser(repository, bus))
			auth.POST("/login", action.HandleLoginUser(repository))
			auth.PATCH("/refresh", action.HandleRefreshTokens(repository))
			// router.POST("/request-reset", auth.HandleRequestPasswordReset)
			// router.POST("/reset-password", auth.HandleResetPassword)
		}
	}
}
