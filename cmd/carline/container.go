package main

import (
	"github.com/pascalallen/carline/internal/carline/domain/permission"
	"github.com/pascalallen/carline/internal/carline/domain/role"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/domain/student"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/database"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/rabbitmq/amqp091-go"
)

type Container struct {
	DatabaseSession        database.Session
	PermissionRepository   permission.Repository
	RoleRepository         role.Repository
	UserRepository         user.Repository
	SchoolRepository       school.Repository
	StudentRepository      student.Repository
	DatabaseSeeder         database.Seeder
	MessageQueueConnection *amqp091.Connection
	CommandBus             messaging.CommandBus
	QueryBus               messaging.QueryBus
	EventDispatcher        messaging.EventDispatcher
}

func NewContainer(
	dbSession database.Session,
	permissionRepo permission.Repository,
	roleRepo role.Repository,
	userRepo user.Repository,
	schoolRepo school.Repository,
	studentRepo student.Repository,
	dbSeeder database.Seeder,
	mqConn *amqp091.Connection,
	commandBus messaging.CommandBus,
	queryBus messaging.QueryBus,
	eventDispatcher messaging.EventDispatcher,
) Container {
	return Container{
		DatabaseSession:        dbSession,
		PermissionRepository:   permissionRepo,
		RoleRepository:         roleRepo,
		UserRepository:         userRepo,
		SchoolRepository:       schoolRepo,
		StudentRepository:      studentRepo,
		DatabaseSeeder:         dbSeeder,
		MessageQueueConnection: mqConn,
		CommandBus:             commandBus,
		QueryBus:               queryBus,
		EventDispatcher:        eventDispatcher,
	}
}
