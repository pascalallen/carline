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
	if os.Getenv("APP_ENV") == "production" {
		sendgridClient := mail.NewSendGridMailClient()
		return mail.NewSendGridMailService(sendgridClient)
	}

	host := os.Getenv("MAILTRAP_HOST")
	port := os.Getenv("MAILTRAP_PORT")
	username := os.Getenv("MAILTRAP_USERNAME")
	password := os.Getenv("MAILTRAP_PASSWORD")

	return mail.NewMailtrapMailService(host, port, username, password)
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
		mail.NewSendGridMailClient,
		provideMailService,
		security_token.NewService,
	)
	return Container{}
}
