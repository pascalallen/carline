//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/pascalallen/carline/internal/carline/infrastructure/database"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/repository"
)

func InitializeContainer() Container {
	wire.Build(
		NewContainer,
		database.NewGormSession,
		repository.NewGormPermissionRepository,
		repository.NewGormRoleRepository,
		repository.NewGormUserRepository,
		repository.NewGormSchoolRepository,
		repository.NewGormStudentRepository,
		database.NewPostgresSeeder,
		messaging.NewRabbitMQConnection,
		messaging.NewRabbitMqCommandBus,
		messaging.NewSynchronousQueryBus,
		messaging.NewRabbitMqEventDispatcher,
	)
	return Container{}
}
