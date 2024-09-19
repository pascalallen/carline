package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/command_handler"
	"github.com/pascalallen/carline/internal/carline/application/event"
	"github.com/pascalallen/carline/internal/carline/application/listener"
	"github.com/pascalallen/carline/internal/carline/domain/permission"
	"github.com/pascalallen/carline/internal/carline/domain/role"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/domain/student"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/routes"
	"os"
)

func main() {
	container := InitializeContainer()
	defer container.MessageQueueConnection.Close()

	configureDatabase(container)

	runConsumers(container)

	configureServer(container)
}

func configureDatabase(container Container) {
	container.DatabaseSession.AutoMigrate(&permission.Permission{}, &role.Role{}, &user.User{}, &school.School{}, &student.Student{})
	container.DatabaseSeeder.Seed()
}

func runConsumers(container Container) {
	commandBus := container.CommandBus
	eventDispatcher := container.EventDispatcher
	userRepository := container.UserRepository

	// command registry
	commandBus.RegisterHandler(command.RegisterUser{}.CommandName(), command_handler.RegisterUserHandler{UserRepository: userRepository, EventDispatcher: eventDispatcher})
	commandBus.RegisterHandler(command.UpdateUser{}.CommandName(), command_handler.UpdateUserHandler{})
	commandBus.RegisterHandler(command.SendWelcomeEmail{}.CommandName(), command_handler.SendWelcomeEmailHandler{EventDispatcher: eventDispatcher})

	// event registry
	eventDispatcher.RegisterListener(event.UserRegistered{}.EventName(), listener.UserRegistration{CommandBus: commandBus})

	go commandBus.StartConsuming()
	go eventDispatcher.StartConsuming()
}

func configureServer(container Container) {
	commandBus := container.CommandBus
	userRepository := container.UserRepository

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := routes.NewRouter()

	router.Config()
	router.Fileserver()
	router.Default()
	router.Auth(userRepository, commandBus)
	router.Temp(userRepository)
	router.Serve(":9990")
}
