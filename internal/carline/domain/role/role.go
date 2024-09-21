package role

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/permission"
	_type "github.com/pascalallen/carline/internal/carline/infrastructure/database/type"
	"time"
)

type Role struct {
	Id          _type.GormUlid          `json:"id" gorm:"primaryKey;size:26;not null"`
	Name        string                  `json:"name"`
	Permissions []permission.Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions"`
	CreatedAt   time.Time               `json:"created_at"`
	ModifiedAt  *time.Time              `json:"modified_at,omitempty"`
}

func Define(id ulid.ULID, name string) *Role {
	createdAt := time.Now()

	return &Role{
		Id:        _type.GormUlid(id),
		Name:      name,
		CreatedAt: createdAt,
	}
}

func (r *Role) UpdateName(name string) {
	r.Name = name
	now := time.Now()
	r.ModifiedAt = &now
}

func (r *Role) AddPermission(permission permission.Permission) {
	for _, p := range r.Permissions {
		if p.Id == permission.Id {
			return
		}
	}

	r.Permissions = append(r.Permissions, permission)
	now := time.Now()
	r.ModifiedAt = &now
}

func (r *Role) RemovePermission(permission permission.Permission) {
	for i, p := range r.Permissions {
		if p.Id == permission.Id {
			r.Permissions[i] = r.Permissions[len(r.Permissions)-1]
		}
	}

	r.Permissions = r.Permissions[:len(r.Permissions)-1]
}

func (r *Role) HasPermission(name string) bool {
	for _, p := range r.Permissions {
		if p.Name == name {
			return true
		}
	}

	return false
}
