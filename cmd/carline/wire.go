//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/storage"
)

func InitializeContainer() Container {
	wire.Build(
		NewContainer,
		storage.NewEventStoreDb,
		messaging.NewRabbitMQConnection,
		messaging.NewRabbitMqCommandBus,
		messaging.NewSynchronousQueryBus,
	)
	return Container{}
}
