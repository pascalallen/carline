package query_handler

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/student"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type ListStudentsHandler struct {
	StudentRepository student.Repository
}

func (h ListStudentsHandler) Handle(qry messaging.Query) (any, error) {
	q, ok := qry.(query.ListStudents)
	if !ok {
		return nil, fmt.Errorf("invalid query type passed to ListStudentsHandler: %v", qry)
	}

	s, err := h.StudentRepository.GetAll(q.SchoolId, q.IncludeDeleted)
	if err != nil {
		return nil, fmt.Errorf("error attempting to list Schools from database: %s", err)
	}

	return s, nil
}
