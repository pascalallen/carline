package security_token

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/crypto"
	"github.com/pascalallen/carline/internal/carline/domain/user"
)

type Repository interface {
	GetById(id ulid.ULID) (*SecurityToken, error)
	GetByCrypto(crypto crypto.Crypto) (*SecurityToken, error)
	GetAllForUser(user user.User) (*[]SecurityToken, error)
	ClearAllForUser(user user.User) error
	Add(securityToken *SecurityToken) error
	Remove(securityToken *SecurityToken) error
}
