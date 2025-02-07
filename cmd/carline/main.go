package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/command_handler"
	"github.com/pascalallen/carline/internal/carline/application/event"
	"github.com/pascalallen/carline/internal/carline/application/listener"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/application/query_handler"
	"github.com/pascalallen/carline/internal/carline/infrastructure/container"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/routes"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	container := container.InitializeContainer()
	defer container.MessageQueueConnection.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		runConsumers(container)
		configureServer(container)
	}()

	<-stop
	fmt.Println("\nShutting down gracefully...")
}

func runConsumers(container container.Container) {
	setupCommandHandlers(container.CommandBus, container)
	setupEventListeners(container.EventDispatcher, container)
	setupQueryHandlers(container.QueryBus, container)

	go container.CommandBus.StartConsuming()
	go container.EventDispatcher.StartConsuming()
}

func setupCommandHandlers(commandBus messaging.CommandBus, container container.Container) {
	commandBus.RegisterHandler(command.RegisterUser{}.CommandName(), command_handler.RegisterUserHandler{
		UserRepository:       container.UserRepository,
		RoleRepository:       container.RoleRepository,
		SecurityTokenService: container.SecurityTokenService,
		EventDispatcher:      container.EventDispatcher,
	})
	commandBus.RegisterHandler(command.UpdateUser{}.CommandName(), command_handler.UpdateUserHandler{})
	commandBus.RegisterHandler(command.SendWelcomeEmail{}.CommandName(), command_handler.SendWelcomeEmailHandler{
		SecurityTokenService: container.SecurityTokenService,
		EventDispatcher:      container.EventDispatcher,
		MailService:          container.MailService,
	})
	commandBus.RegisterHandler(command.CreateSchool{}.CommandName(), command_handler.CreateSchoolHandler{
		SchoolRepository: container.SchoolRepository,
	})
	commandBus.RegisterHandler(command.DeleteSchool{}.CommandName(), command_handler.DeleteSchoolHandler{
		SchoolRepository: container.SchoolRepository,
	})
	commandBus.RegisterHandler(command.ImportStudents{}.CommandName(), command_handler.ImportStudentsHandler{
		SchoolRepository:  container.SchoolRepository,
		StudentRepository: container.StudentRepository,
		DatabaseSession:   container.DatabaseSession,
	})
	commandBus.RegisterHandler(command.DeleteStudent{}.CommandName(), command_handler.DeleteStudentHandler{
		StudentRepository: container.StudentRepository,
	})
}

func setupEventListeners(eventDispatcher messaging.EventDispatcher, container container.Container) {
	eventDispatcher.RegisterListener(event.UserRegistered{}.EventName(), listener.UserRegistration{
		CommandBus: container.CommandBus,
	})
}

func setupQueryHandlers(queryBus messaging.QueryBus, container container.Container) {
	queryBus.RegisterHandler(query.GetUserById{}.QueryName(), query_handler.GetUserByIdHandler{
		UserRepository: container.UserRepository,
	})
	queryBus.RegisterHandler(query.GetUserByEmailAddress{}.QueryName(), query_handler.GetUserByEmailAddressHandler{
		UserRepository: container.UserRepository,
	})
	queryBus.RegisterHandler(query.GetSchoolByName{}.QueryName(), query_handler.GetSchoolByNameHandler{
		SchoolRepository: container.SchoolRepository,
	})
	queryBus.RegisterHandler(query.ListSchools{}.QueryName(), query_handler.ListSchoolsHandler{
		SchoolRepository: container.SchoolRepository,
	})
	queryBus.RegisterHandler(query.GetSchoolByIdAndUserId{}.QueryName(), query_handler.GetSchoolByIdAndUserIdHandler{
		SchoolRepository: container.SchoolRepository,
	})
	queryBus.RegisterHandler(query.ListStudents{}.QueryName(), query_handler.ListStudentsHandler{
		StudentRepository: container.StudentRepository,
	})
	queryBus.RegisterHandler(query.GetStudentById{}.QueryName(), query_handler.GetStudentByIdHandler{
		StudentRepository: container.StudentRepository,
	})
}

func configureServer(container container.Container) {
	queryBus := container.QueryBus
	commandBus := container.CommandBus
	securityTokenService := container.SecurityTokenService
	userRepository := container.UserRepository

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := routes.NewRouter()

	router.Config()
	router.Fileserver()
	router.Default()
	router.Auth(queryBus, commandBus, securityTokenService, userRepository)
	router.Schools(queryBus, commandBus)
	router.Temp(queryBus)
	router.Serve(":9990")
}
