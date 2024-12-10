//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/pascalallen/carline/internal/carline/infrastructure/database"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/repository"
	"github.com/pascalallen/carline/internal/carline/infrastructure/service"
)

func InitializeContainer() Container {
	wire.Build(
		NewContainer,
		database.NewPostgresSession,
		repository.NewPostgresPermissionRepository,
		repository.NewPostgresRoleRepository,
		repository.NewGormUserRepository,
		repository.NewPostgresSchoolRepository,
		repository.NewPostgresStudentRepository,
		messaging.NewRabbitMQConnection,
		messaging.NewRabbitMqCommandBus,
		messaging.NewSynchronousQueryBus,
		messaging.NewRabbitMqEventDispatcher,
		service.NewSendGridMailService,
	)
	return Container{}
}
