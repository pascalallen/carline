package command_handler

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type AddSchoolHandler struct {
	SchoolRepository school.Repository
}

func (h AddSchoolHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.AddSchool)
	if !ok {
		return fmt.Errorf("invalid command type passed to AddSchoolHandler: %v", cmd)
	}

	s := school.Create(c.Id, c.Name)

	err := h.SchoolRepository.Add(s)
	if err != nil {
		return fmt.Errorf("school creation failed: %s", err)
	}

	return nil
}
