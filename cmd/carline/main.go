package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/command_handler"
	"github.com/pascalallen/carline/internal/carline/application/projection"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/application/query_handler"
	"github.com/pascalallen/carline/internal/carline/infrastructure/routes"
	"log"
	"os"
)

func main() {
	container := InitializeContainer()
	defer container.MessageQueueConnection.Close()

	setupProjections(container)

	runConsumers(container)

	configureServer(container)
}

func setupProjections(container Container) {
	eventStore := container.EventStore

	// projection registry
	err := eventStore.RegisterProjection(projection.UserEmailAddresses{})
	if err != nil {
		exitOnError(err)
	}
}

func runConsumers(container Container) {
	commandBus := container.CommandBus
	queryBus := container.QueryBus
	eventStore := container.EventStore

	// command registry
	commandBus.RegisterHandler(command.UpdateUserEmailAddress{}.CommandName(), command_handler.UpdateUserEmailAddressHandler{EventStore: eventStore})

	// query registry
	queryBus.RegisterHandler(query.GetUserById{}.QueryName(), query_handler.GetUserByIdHandler{EventStore: eventStore})
	queryBus.RegisterHandler(query.GetUserByEmailAddress{}.QueryName(), query_handler.GetUserByEmailAddressHandler{EventStore: eventStore})

	go commandBus.StartConsuming()
}

func configureServer(container Container) {
	eventStore := container.EventStore
	queryBus := container.QueryBus

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := routes.NewRouter()

	router.Config()
	router.Fileserver()
	router.Default()
	router.Auth(queryBus, eventStore)
	router.Temp(queryBus)
	router.Serve(":9990")
}

func exitOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
