package main

import (
	"database/sql"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/pascalallen/carline/internal/carline/domain/mail"
	"github.com/pascalallen/carline/internal/carline/domain/permission"
	"github.com/pascalallen/carline/internal/carline/domain/role"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/domain/security_token"
	"github.com/pascalallen/carline/internal/carline/domain/student"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/rabbitmq/amqp091-go"
)

type Container struct {
	DatabaseSession         *sql.DB
	PermissionRepository    permission.Repository
	RoleRepository          role.Repository
	UserRepository          user.Repository
	SecurityTokenRepository security_token.Repository
	SchoolRepository        school.Repository
	StudentRepository       student.Repository
	MessageQueueConnection  *amqp091.Connection
	CommandBus              messaging.CommandBus
	QueryBus                messaging.QueryBus
	EventDispatcher         messaging.EventDispatcher
	MailClient              *mailgun.MailgunImpl
	MailService             mail.Service
	SecurityTokenService    security_token.Service
}

func NewContainer(
	dbSession *sql.DB,
	permissionRepo permission.Repository,
	roleRepo role.Repository,
	userRepo user.Repository,
	securityTokenRepo security_token.Repository,
	schoolRepo school.Repository,
	studentRepo student.Repository,
	mqConn *amqp091.Connection,
	commandBus messaging.CommandBus,
	queryBus messaging.QueryBus,
	eventDispatcher messaging.EventDispatcher,
	mailClient *mailgun.MailgunImpl,
	mailService mail.Service,
	securityTokenService security_token.Service,
) Container {
	return Container{
		DatabaseSession:         dbSession,
		PermissionRepository:    permissionRepo,
		RoleRepository:          roleRepo,
		UserRepository:          userRepo,
		SecurityTokenRepository: securityTokenRepo,
		SchoolRepository:        schoolRepo,
		StudentRepository:       studentRepo,
		MessageQueueConnection:  mqConn,
		CommandBus:              commandBus,
		QueryBus:                queryBus,
		EventDispatcher:         eventDispatcher,
		MailClient:              mailClient,
		MailService:             mailService,
		SecurityTokenService:    securityTokenService,
	}
}
