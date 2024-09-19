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
	PermissionRepository   permission.PermissionRepository
	RoleRepository         role.RoleRepository
	UserRepository         user.UserRepository
	SchoolRepository       school.SchoolRepository
	StudentRepository      student.StudentRepository
	DatabaseSeeder         database.Seeder
	MessageQueueConnection *amqp091.Connection
	CommandBus             messaging.CommandBus
	EventDispatcher        messaging.EventDispatcher
}

func NewContainer(
	dbSession database.Session,
	permissionRepo permission.PermissionRepository,
	roleRepo role.RoleRepository,
	userRepo user.UserRepository,
	schoolRepo school.SchoolRepository,
	studentRepo student.StudentRepository,
	dbSeeder database.Seeder,
	mqConn *amqp091.Connection,
	commandBus messaging.CommandBus,
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
		EventDispatcher:        eventDispatcher,
	}
}
