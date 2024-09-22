package school

import "github.com/oklog/ulid/v2"

type Repository interface {
	GetById(id ulid.ULID) (*School, error)
	GetByName(name string) (*School, error)
	GetAll(includeDeleted bool) (*[]School, error)
	Add(school *School) error
	Remove(school *School) error
}
