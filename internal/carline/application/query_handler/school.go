package query_handler

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type GetSchoolByNameHandler struct {
	SchoolRepository school.Repository
}

func (h GetSchoolByNameHandler) Handle(qry messaging.Query) (any, error) {
	q, ok := qry.(query.GetSchoolByName)
	if !ok {
		return nil, fmt.Errorf("invalid query type passed to GetSchoolByNameHandler: %v", qry)
	}

	s, err := h.SchoolRepository.GetByName(q.Name)
	if err != nil {
		return nil, fmt.Errorf("error attempting to retrieve School from database: %s", err)
	}

	return s, nil
}

type ListSchoolsHandler struct {
	SchoolRepository school.Repository
}

func (h ListSchoolsHandler) Handle(qry messaging.Query) (any, error) {
	q, ok := qry.(query.ListSchools)
	if !ok {
		return nil, fmt.Errorf("invalid query type passed to ListSchoolsHandler: %v", qry)
	}

	s, err := h.SchoolRepository.GetAll(q.UserId)
	if err != nil {
		return nil, fmt.Errorf("error attempting to list Schools from database: %s", err)
	}

	return s, nil
}

type GetSchoolByIdAndUserIdHandler struct {
	SchoolRepository school.Repository
}

func (h GetSchoolByIdAndUserIdHandler) Handle(qry messaging.Query) (any, error) {
	q, ok := qry.(query.GetSchoolByIdAndUserId)
	if !ok {
		return nil, fmt.Errorf("invalid query type passed to GetSchoolByIdAndUserIdHandler: %v", qry)
	}

	s, err := h.SchoolRepository.GetByIdAndUserId(q.UserId, q.Id)
	if err != nil {
		return nil, fmt.Errorf("error attempting to fetch School from database: %s", err)
	}

	return s, nil
}
