package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/command_handler"
	"github.com/pascalallen/carline/internal/carline/application/event"
	"github.com/pascalallen/carline/internal/carline/application/listener"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/application/query_handler"
	"github.com/pascalallen/carline/internal/carline/infrastructure/routes"
	"os"
)

func main() {
	container := InitializeContainer()
	defer container.MessageQueueConnection.Close()

	runConsumers(container)

	configureServer(container)
}

func runConsumers(container Container) {
	commandBus := container.CommandBus
	queryBus := container.QueryBus
	eventDispatcher := container.EventDispatcher
	userRepository := container.UserRepository
	schoolRepository := container.SchoolRepository
	studentRepository := container.StudentRepository
	databaseSession := container.DatabaseSession

	// command registry
	commandBus.RegisterHandler(command.RegisterUser{}.CommandName(), command_handler.RegisterUserHandler{UserRepository: userRepository, EventDispatcher: eventDispatcher})
	commandBus.RegisterHandler(command.UpdateUser{}.CommandName(), command_handler.UpdateUserHandler{})
	commandBus.RegisterHandler(command.SendWelcomeEmail{}.CommandName(), command_handler.SendWelcomeEmailHandler{EventDispatcher: eventDispatcher})
	commandBus.RegisterHandler(command.CreateSchool{}.CommandName(), command_handler.CreateSchoolHandler{SchoolRepository: schoolRepository})
	commandBus.RegisterHandler(command.DeleteSchool{}.CommandName(), command_handler.DeleteSchoolHandler{SchoolRepository: schoolRepository})
	commandBus.RegisterHandler(command.ImportStudents{}.CommandName(), command_handler.ImportStudentsHandler{SchoolRepository: schoolRepository, StudentRepository: studentRepository, DatabaseSession: databaseSession})
	commandBus.RegisterHandler(command.DeleteStudent{}.CommandName(), command_handler.DeleteStudentHandler{StudentRepository: studentRepository})

	// event registry
	eventDispatcher.RegisterListener(event.UserRegistered{}.EventName(), listener.UserRegistration{CommandBus: commandBus})

	// query registry
	queryBus.RegisterHandler(query.GetUserById{}.QueryName(), query_handler.GetUserByIdHandler{UserRepository: userRepository})
	queryBus.RegisterHandler(query.GetUserByEmailAddress{}.QueryName(), query_handler.GetUserByEmailAddressHandler{UserRepository: userRepository})
	queryBus.RegisterHandler(query.GetSchoolByName{}.QueryName(), query_handler.GetSchoolByNameHandler{SchoolRepository: schoolRepository})
	queryBus.RegisterHandler(query.ListSchools{}.QueryName(), query_handler.ListSchoolsHandler{SchoolRepository: schoolRepository})
	queryBus.RegisterHandler(query.GetSchoolByIdAndUserId{}.QueryName(), query_handler.GetSchoolByIdAndUserIdHandler{SchoolRepository: schoolRepository})
	queryBus.RegisterHandler(query.ListStudents{}.QueryName(), query_handler.ListStudentsHandler{StudentRepository: studentRepository})
	queryBus.RegisterHandler(query.GetStudentById{}.QueryName(), query_handler.GetStudentByIdHandler{StudentRepository: studentRepository})

	go commandBus.StartConsuming()
	go eventDispatcher.StartConsuming()
}

func configureServer(container Container) {
	queryBus := container.QueryBus
	commandBus := container.CommandBus

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := routes.NewRouter()

	router.Config()
	router.Fileserver()
	router.Default()
	router.Auth(queryBus, commandBus)
	router.Schools(queryBus, commandBus)
	router.Temp(queryBus)
	router.Serve(":9990")
}
