package command_handler

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/command"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type CreateSchoolHandler struct {
	SchoolRepository school.Repository
}

func (h CreateSchoolHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.CreateSchool)
	if !ok {
		return fmt.Errorf("invalid command type passed to CreateSchoolHandler: %v", cmd)
	}

	s := school.Create(c.Id, c.Name)

	err := h.SchoolRepository.Add(s)
	if err != nil {
		return fmt.Errorf("school creation failed: %s", err)
	}

	return nil
}

type DeleteSchoolHandler struct {
	SchoolRepository school.Repository
}

func (h DeleteSchoolHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.DeleteSchool)
	if !ok {
		return fmt.Errorf("invalid command type passed to DeleteSchoolHandler: %v", cmd)
	}

	s, err := h.SchoolRepository.GetById(c.Id)
	if err != nil {
		return fmt.Errorf("school not found: %s", c.Id)
	}

	err = h.SchoolRepository.Remove(s)
	if err != nil {
		return fmt.Errorf("school removal failed: %s", err)
	}

	return nil
}
