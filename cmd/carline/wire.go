//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/pascalallen/carline/internal/carline/domain/security_token"
	"github.com/pascalallen/carline/internal/carline/infrastructure/database"
	"github.com/pascalallen/carline/internal/carline/infrastructure/mail"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/repository"
)

func InitializeContainer() Container {
	wire.Build(
		NewContainer,
		database.NewPostgresSession,
		repository.NewPostgresPermissionRepository,
		repository.NewPostgresRoleRepository,
		repository.NewPostgresUserRepository,
		repository.NewPostgresSecurityTokenRepository,
		repository.NewPostgresSchoolRepository,
		repository.NewPostgresStudentRepository,
		messaging.NewRabbitMQConnection,
		messaging.NewRabbitMqCommandBus,
		messaging.NewSynchronousQueryBus,
		messaging.NewRabbitMqEventDispatcher,
		mail.NewSendGridMailClient,
		mail.NewSendGridMailService,
		security_token.NewService,
	)
	return Container{}
}
