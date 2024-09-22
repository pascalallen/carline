package query_handler

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/application/query"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"github.com/pascalallen/carline/internal/carline/infrastructure/messaging"
)

type GetUserByIdHandler struct {
	UserRepository user.Repository
}

func (h GetUserByIdHandler) Handle(qry messaging.Query) (any, error) {
	q, ok := qry.(query.GetUserById)
	if !ok {
		return nil, fmt.Errorf("invalid query type passed to GetUserByIdHandler: %v", qry)
	}

	u, err := h.UserRepository.GetById(q.Id)
	if err != nil {
		return nil, fmt.Errorf("error attempting to retrieve user from database: %s", err)
	}

	return u, nil
}

type GetUserByEmailAddressHandler struct {
	UserRepository user.Repository
}

func (h GetUserByEmailAddressHandler) Handle(qry messaging.Query) (any, error) {
	q, ok := qry.(query.GetUserByEmailAddress)
	if !ok {
		return nil, fmt.Errorf("invalid query type passed to GetUserByEmailAddressHandler: %v", qry)
	}

	u, err := h.UserRepository.GetByEmailAddress(q.EmailAddress)
	if err != nil {
		return nil, fmt.Errorf("error attempting to retrieve user from database: %s", err)
	}

	return u, nil
}
