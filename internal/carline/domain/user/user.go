package user

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/password"
	"github.com/pascalallen/carline/internal/carline/domain/permission"
	"github.com/pascalallen/carline/internal/carline/domain/role"
	"github.com/pascalallen/carline/internal/carline/domain/school"
	"time"
)

type User struct {
	Id           ulid.ULID             `json:"id"`
	FirstName    string                `json:"first_name"`
	LastName     string                `json:"last_name"`
	EmailAddress string                `json:"email_address"`
	PasswordHash password.PasswordHash `json:"-"`
	Schools      []school.School       `json:"schools"`
	Roles        []role.Role           `json:"roles"`
	CreatedAt    time.Time             `json:"created_at"`
	ModifiedAt   *time.Time            `json:"modified_at,omitempty"`
}

func Register(id ulid.ULID, firstName string, lastName string, emailAddress string) *User {
	createdAt := time.Now()

	return &User{
		Id:           id,
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		CreatedAt:    createdAt,
	}
}

func (u *User) UpdateFirstName(firstName string) {
	u.FirstName = firstName
	now := time.Now()
	u.ModifiedAt = &now
}

func (u *User) UpdateLastName(lastName string) {
	u.LastName = lastName
	now := time.Now()
	u.ModifiedAt = &now
}

func (u *User) UpdateEmailAddress(emailAddress string) {
	u.EmailAddress = emailAddress
	now := time.Now()
	u.ModifiedAt = &now
}

func (u *User) SetPasswordHash(passwordHash password.PasswordHash) {
	u.PasswordHash = passwordHash
	now := time.Now()
	u.ModifiedAt = &now
}

func (u *User) AddSchool(school school.School) {
	for _, s := range u.Schools {
		if s.Id == school.Id {
			return
		}
	}

	u.Schools = append(u.Schools, school)
	now := time.Now()
	u.ModifiedAt = &now
}

func (u *User) RemoveSchool(school school.School) {
	for i, s := range u.Schools {
		if s.Id == school.Id {
			u.Schools[i] = u.Schools[len(u.Schools)-1]
			u.Schools = u.Schools[:len(u.Schools)-1]
			now := time.Now()
			u.ModifiedAt = &now
			return
		}
	}
}

func (u *User) HasSchool(schoolId ulid.ULID) bool {
	for _, s := range u.Schools {
		if s.Id == schoolId {
			return true
		}
	}

	return false
}

func (u *User) AddRole(role role.Role) {
	for _, r := range u.Roles {
		if r.Id == role.Id {
			return
		}
	}

	u.Roles = append(u.Roles, role)
	now := time.Now()
	u.ModifiedAt = &now
}

func (u *User) RemoveRole(role role.Role) {
	for i, r := range u.Roles {
		if r.Id == role.Id {
			u.Roles[i] = u.Roles[len(u.Roles)-1]
		}
	}

	u.Roles = u.Roles[:len(u.Roles)-1]
}

func (u *User) HasRole(name string) bool {
	for _, r := range u.Roles {
		if r.Name == name {
			return true
		}
	}

	return false
}

func (u *User) Permissions() []permission.Permission {
	var permissions []permission.Permission
	for _, r := range u.Roles {
		permissions = append(permissions, r.Permissions...)
	}

	return permissions
}

func (u *User) HasPermission(name string) bool {
	for _, p := range u.Permissions() {
		if p.Name == name {
			return true
		}
	}

	return false
}
