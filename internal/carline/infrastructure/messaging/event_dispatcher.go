package messaging

import (
	"context"
	"encoding/json"
	"github.com/pascalallen/carline/internal/carline/application/event"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Event interface {
	EventName() string
}

type Listener interface {
	Handle(event Event) error
}

type EventDispatcher interface {
	RegisterListener(eventType string, listener Listener)
	StartConsuming()
	Dispatch(evt Event)
}

type RabbitMqEventDispatcher struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	listeners  map[string]Listener
}

const exchangeName = "events"

func NewRabbitMqEventDispatcher(conn *amqp091.Connection) EventDispatcher {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open server channel for event dispatcher: %s", err)
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare exchange: %s", err)
	}

	return &RabbitMqEventDispatcher{
		connection: conn,
		channel:    ch,
		listeners:  make(map[string]Listener),
	}
}

func (e *RabbitMqEventDispatcher) RegisterListener(eventType string, listener Listener) {
	e.listeners[eventType] = listener
}

func (e *RabbitMqEventDispatcher) StartConsuming() {
	queueName := e.setupQueue()
	msgs := e.messages(queueName)

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			e.processEvent(msg)
		}
	}()

	<-forever
}

func (e *RabbitMqEventDispatcher) Dispatch(evt Event) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	b, err := json.Marshal(evt)
	if err != nil {
		log.Fatalf("failed to JSON encode event: %s", err)
	}

	if e.channel.IsClosed() {
		ch, err := e.connection.Channel()
		if err != nil {
			log.Fatalf("failed to recreate server channel for event dispatcher: %s", err)
		}

		e.channel = ch
	}

	err = e.channel.PublishWithContext(
		ctx,
		exchangeName,
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        b,
			Type:        evt.EventName(),
		},
	)
	if err != nil {
		log.Fatalf("failed to dispatch event: %s", err)
	}
}

func (e *RabbitMqEventDispatcher) setupQueue() string {
	q, err := e.channel.QueueDeclare(
		exchangeName,
		true,  // durable
		false, // auto-delete
		false, // not exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare queue: %s", err)
	}

	err = e.channel.QueueBind(
		q.Name,
		"",
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to bind to queue: %s", err)
	}

	return q.Name
}

func (e *RabbitMqEventDispatcher) messages(queueName string) <-chan amqp091.Delivery {
	// No need to declare the exchange again here
	d, err := e.channel.Consume(
		queueName,
		"",
		false, // auto-ack disabled
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to consume event messages: %s", err)
	}

	return d
}

func (e *RabbitMqEventDispatcher) processEvent(msg amqp091.Delivery) {
	switch msg.Type {
	case event.UserRegistered{}.EventName(): // Handle UserRegistered event
		evt := &event.UserRegistered{}
		err := json.Unmarshal(msg.Body, evt)
		if err != nil {
			log.Println("Failed to unmarshal event:", err)
			_ = msg.Nack(false, true) // Requeue the message for retry
			return
		}
		e.handleEvent(evt, msg)

	case event.UserUpdated{}.EventName(): // Handle UserUpdated event
		evt := &event.UserUpdated{}
		err := json.Unmarshal(msg.Body, evt)
		if err != nil {
			log.Println("Failed to unmarshal event:", err)
			_ = msg.Nack(false, true) // Requeue the message for retry
			return
		}
		e.handleEvent(evt, msg)

	case event.WelcomeEmailSent{}.EventName(): // Handle WelcomeEmailSent event
		evt := &event.WelcomeEmailSent{}
		err := json.Unmarshal(msg.Body, evt)
		if err != nil {
			log.Println("Failed to unmarshal event:", err)
			_ = msg.Nack(false, true) // Requeue the message for retry
			return
		}
		e.handleEvent(evt, msg)

	default:
		log.Printf("Unknown event received: %s", msg.Type)
		_ = msg.Ack(false) // Acknowledge unknown events and drop them
	}
}

func (e *RabbitMqEventDispatcher) handleEvent(evt Event, msg amqp091.Delivery) {
	listener, found := e.listeners[evt.EventName()]
	if !found {
		log.Printf("No listener registered for event type: %s", evt.EventName())
		_ = msg.Ack(false) // Acknowledge the message to drop it
		return
	}

	err := listener.Handle(evt)
	if err != nil {
		log.Printf("Error calling listener for event %s: %s", evt.EventName(), err)
		_ = msg.Nack(false, true) // Requeue the message for retry
		return
	}

	err = msg.Ack(false) // Acknowledge successful processing
	if err != nil {
		log.Printf("Error acknowledging event message: %s", err)
		return
	}
}
