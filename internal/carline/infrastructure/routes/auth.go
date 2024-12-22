package routes

import (
	"github.com/pascalallen/carline/internal/carline/application/http/action/auth"
	"github.com/pascalallen/carline/internal/carline/domain/security_token"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

func (r Router) Auth(queryBus messaging.QueryBus, commandBus messaging.CommandBus, securityTokenService security_token.Service, userRepo user.Repository) {
	v := r.engine.Group(v1)
	{
		a := v.Group("/auth")
		{
			a.POST("/register", auth.HandleRegisterUser(queryBus, commandBus))
			a.POST("/activate", auth.HandleActivateUser(queryBus, commandBus, securityTokenService, userRepo))
			a.POST("/login", auth.HandleLoginUser(queryBus))
			a.PATCH("/refresh", auth.HandleRefreshTokens(queryBus))
			// router.POST("/request-reset", auth.HandleRequestPasswordReset)
			// router.POST("/reset-password", auth.HandleResetPassword)
		}
	}
}
