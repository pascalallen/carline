package school

import (
	"github.com/oklog/ulid/v2"
)

type Repository interface {
	GetByIdAndUserId(userId ulid.ULID, id ulid.ULID) (*School, error)
	GetById(id ulid.ULID) (*School, error)
	GetByName(name string) (*School, error)
	GetAll(userId ulid.ULID) (*[]School, error)
	Add(school *School, userId ulid.ULID) error
	Remove(school *School, userId ulid.ULID) error
}
