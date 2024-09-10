package main

import (
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/storage"
	"github.com/rabbitmq/amqp091-go"
)

type Container struct {
	EventStore             storage.EventStore
	MessageQueueConnection *amqp091.Connection
	CommandBus             messaging.CommandBus
	QueryBus               messaging.QueryBus
}

func NewContainer(
	eventStore storage.EventStore,
	mqConn *amqp091.Connection,
	commandBus messaging.CommandBus,
	queryBus messaging.QueryBus,
) Container {
	return Container{
		EventStore:             eventStore,
		MessageQueueConnection: mqConn,
		CommandBus:             commandBus,
		QueryBus:               queryBus,
	}
}
