package routes

import (
	"github.com/pascalallen/carline/internal/carline/application/http/action/auth"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

func (r Router) Auth(repository user.Repository, bus messaging.CommandBus) {
	v := r.engine.Group(v1)
	{
		a := v.Group("/auth")
		{
			a.POST("/register", auth.HandleRegisterUser(repository, bus))
			a.POST("/login", auth.HandleLoginUser(repository))
			a.PATCH("/refresh", auth.HandleRefreshTokens(repository))
			// router.POST("/request-reset", auth.HandleRequestPasswordReset)
			// router.POST("/reset-password", auth.HandleResetPassword)
		}
	}
}
