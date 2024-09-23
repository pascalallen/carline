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

type RemoveSchoolHandler struct {
	SchoolRepository school.Repository
}

func (h RemoveSchoolHandler) Handle(cmd messaging.Command) error {
	c, ok := cmd.(*command.RemoveSchool)
	if !ok {
		return fmt.Errorf("invalid command type passed to RemoveSchoolHandler: %v", cmd)
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
