package container

import (
	"database/sql"
	"github.com/pascalallen/carline/internal/carline/domain/mail"
	"github.com/pascalallen/carline/internal/carline/domain/permission"
	"github.com/pascalallen/carline/internal/carline/domain/role"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/domain/security_token"
	"github.com/pascalallen/carline/internal/carline/domain/student"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"github.com/pascalallen/carline/internal/carline/infrastructure/websocket"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sendgrid/sendgrid-go"
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
	MailClient              *sendgrid.Client
	MailService             mail.Service
	SecurityTokenService    security_token.Service
	WebsocketHub            *websocket.Hub
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
	mailClient *sendgrid.Client,
	mailService mail.Service,
	securityTokenService security_token.Service,
	websocketHub *websocket.Hub,
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
		WebsocketHub:            websocketHub,
	}
}
