package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Command interface {
	CommandName() string
}

type CommandHandler interface {
	Handle(command Command) error
}

type CommandBus interface {
	RegisterHandler(commandType string, handler CommandHandler)
	StartConsuming()
	Execute(cmd Command) error
}

type RabbitMqCommandBus struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	handlers   map[string]CommandHandler
}

const queueName = "commands"

func NewRabbitMqCommandBus(conn *amqp091.Connection) CommandBus {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open server channel for event dispatcher: %s", err)
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to create or fetch queue: %s\n", err)
	}

	return &RabbitMqCommandBus{
		connection: conn,
		channel:    ch,
		handlers:   make(map[string]CommandHandler),
	}
}

func (bus *RabbitMqCommandBus) RegisterHandler(commandType string, handler CommandHandler) {
	bus.handlers[commandType] = handler
}

func (bus *RabbitMqCommandBus) StartConsuming() {
	msgs, err := bus.messages()
	if err != nil {
		log.Fatalf("failed to consume command messages: %s", err)
	}

	go func() {
		for msg := range msgs {
			if err := bus.processCommand(msg); err != nil {
				log.Printf("failed to process command: %s", err)
			}
		}
	}()

	select {}
}

func (bus *RabbitMqCommandBus) Execute(cmd Command) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	b, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to JSON encode command: %s", err)
	}

	if bus.channel.IsClosed() {
		ch, err := bus.connection.Channel()
		if err != nil {
			log.Fatalf("failed to recreate server channel for event dispatcher: %s", err)
		}

		bus.channel = ch
	}

	err = bus.channel.PublishWithContext(
		ctx,
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "text/plain",
			Body:         b,
			Type:         cmd.CommandName(),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish command: %s", err)
	}

	return nil
}

func (bus *RabbitMqCommandBus) messages() (<-chan amqp091.Delivery, error) {
	err := bus.channel.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to set QoS: %s", err)
	}

	d, err := bus.channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume command messages: %s", err)
	}

	return d, nil
}

func (bus *RabbitMqCommandBus) processCommand(msg amqp091.Delivery) error {
	var cmd Command

	switch msg.Type {
	case command.RegisterUser{}.CommandName():
		cmd = &command.RegisterUser{}
	case command.UpdateUser{}.CommandName():
		cmd = &command.UpdateUser{}
	case command.CreateSchool{}.CommandName():
		cmd = &command.CreateSchool{}
	case command.DeleteSchool{}.CommandName():
		cmd = &command.DeleteSchool{}
	case command.SendWelcomeEmail{}.CommandName():
		cmd = &command.SendWelcomeEmail{}
	case command.ImportStudents{}.CommandName():
		cmd = &command.ImportStudents{}
	case command.DeleteStudent{}.CommandName():
		cmd = &command.DeleteStudent{}
	case command.DismissStudents{}.CommandName():
		cmd = &command.DismissStudents{}
	default:
		return fmt.Errorf("unknown command received: %s", msg.Type)
	}

	err := json.Unmarshal(msg.Body, &cmd)
	if err != nil {
		return fmt.Errorf("failed to unmarshal command: %s", err)
	}

	handler, found := bus.handlers[cmd.CommandName()]
	if !found {
		return fmt.Errorf("no handler registered for command type: %s", cmd.CommandName())
	}

	err = handler.Handle(cmd)
	if err != nil {
		return fmt.Errorf("error calling command handler: %s", err)
	}

	err = msg.Ack(false)
	if err != nil {
		return fmt.Errorf("error acknowledging command message: %s", err)
	}

	return nil
}
