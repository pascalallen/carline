package command_handler

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/application/event"
	"github.com/pascalallen/carline/internal/carline/domain/role"
	"github.com/pascalallen/carline/internal/carline/domain/security_token"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
	"log"
	"time"
)

type RegisterUserHandler struct {
	UserRepository       user.Repository
	RoleRepository       role.Repository
	SecurityTokenService security_token.Service
	EventDispatcher      messaging.EventDispatcher
}

var adminRoles = []string{"ROLE_USER", "ROLE_ADMIN"}

func (h RegisterUserHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.RegisterUser)
	if !ok {
		return fmt.Errorf("invalid command type passed to RegisterUserHandler: %v", cmd)
	}

	u := user.Register(c.Id, c.FirstName, c.LastName, c.EmailAddress)
	if c.Admin {
		for _, roleName := range adminRoles {
			r, err := h.RoleRepository.GetByName(roleName)
			if err != nil {
				return fmt.Errorf("error admin role by name: %s", err)
			}

			if r != nil {
				u.AddRole(*r)
			}
		}

	}

	err := h.UserRepository.Add(u)
	if err != nil {
		return fmt.Errorf("user registration failed: %s", err)
	}

	now := time.Now()
	expiresAt := now.Add(security_token.ActivationDuration)
	token, err := h.SecurityTokenService.GenerateActivationToken(*u, expiresAt)
	if err != nil {
		return fmt.Errorf("error generating activation token: %s", err)
	}

	evt := event.UserRegistered{
		Id:              c.Id,
		FirstName:       c.FirstName,
		LastName:        c.LastName,
		EmailAddress:    c.EmailAddress,
		SecurityTokenId: token.Id,
	}
	h.EventDispatcher.Dispatch(evt)

	return nil
}

type UpdateUserHandler struct{}

func (h UpdateUserHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.UpdateUser)
	if !ok {
		return fmt.Errorf("invalid command type passed to UpdateUserHandler: %v", cmd)
	}

	// TODO
	log.Printf("UpdateUserHandler executed: %v", c)

	return nil
}
