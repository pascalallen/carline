package security_token

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/crypto"
	"time"
)

type SecurityToken struct {
	Id          ulid.ULID         `json:"id"`
	UserId      ulid.ULID         `json:"user_id"`
	Crypto      crypto.Crypto     `json:"crypto"`
	Type        SecurityTokenType `json:"type"`
	GeneratedAt time.Time         `json:"generated_at"`
	ExpiresAt   time.Time         `json:"expires_at"`
	ModifiedAt  *time.Time        `json:"modified_at,omitempty"`
}

func Create(
	id ulid.ULID,
	userId ulid.ULID,
	crypto crypto.Crypto,
	tokenType SecurityTokenType,
	expiresAt time.Time,
) *SecurityToken {
	generatedAt := time.Now()

	return &SecurityToken{
		Id:          id,
		UserId:      userId,
		Crypto:      crypto,
		Type:        tokenType,
		GeneratedAt: generatedAt,
		ExpiresAt:   expiresAt,
	}
}

func GenerateActivation(id ulid.ULID, userId ulid.ULID, expiresAt time.Time) *SecurityToken {
	c := crypto.Generate()
	generatedAt := time.Now()

	return &SecurityToken{
		Id:          id,
		UserId:      userId,
		Crypto:      c,
		Type:        ACTIVATION,
		GeneratedAt: generatedAt,
		ExpiresAt:   expiresAt,
	}
}

func GenerateRefresh(id ulid.ULID, userId ulid.ULID, expiresAt time.Time) *SecurityToken {
	c := crypto.Generate()
	generatedAt := time.Now()

	return &SecurityToken{
		Id:          id,
		UserId:      userId,
		Crypto:      c,
		Type:        REFRESH,
		GeneratedAt: generatedAt,
		ExpiresAt:   expiresAt,
	}
}

func GenerateReset(id ulid.ULID, userId ulid.ULID, expiresAt time.Time) *SecurityToken {
	c := crypto.Generate()
	generatedAt := time.Now()

	return &SecurityToken{
		Id:          id,
		UserId:      userId,
		Crypto:      c,
		Type:        RESET,
		GeneratedAt: generatedAt,
		ExpiresAt:   expiresAt,
	}
}

func (s *SecurityToken) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}
