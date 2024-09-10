package role

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/permission"
	"time"
)

type Role struct {
	Id          ulid.ULID               `json:"id"`
	Name        string                  `json:"name"`
	Permissions []permission.Permission `json:"permissions,omitempty"`
	CreatedAt   time.Time               `json:"created_at"`
	ModifiedAt  time.Time               `json:"modified_at"`
}
