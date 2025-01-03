//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	mail2 "github.com/pascalallen/carline/internal/carline/domain/mail"
	"github.com/pascalallen/carline/internal/carline/domain/security_token"
	"github.com/pascalallen/carline/internal/carline/infrastructure/database"
	"github.com/pascalallen/carline/internal/carline/infrastructure/mail"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/repository"
	"os"
)

func provideMailService() mail2.Service {
	host := os.Getenv("MAILGUN_HOST")
	port := os.Getenv("MAILGUN_PORT")
	username := os.Getenv("MAILGUN_USERNAME")
	password := os.Getenv("MAILGUN_PASSWORD")

	return mail.NewMailgunMailService(host, port, username, password)
}

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
		provideMailService,
		security_token.NewService,
	)
	return Container{}
}
