package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/crypto"
	"github.com/pascalallen/carline/internal/carline/domain/security_token"
	"github.com/pascalallen/carline/internal/carline/domain/user"
)

type PostgresSecurityTokenRepository struct {
	session *sql.DB
}

func (r *PostgresSecurityTokenRepository) GetById(id ulid.ULID) (*security_token.SecurityToken, error) {
	var s security_token.SecurityToken
	var i string
	var uid string
	q := `SELECT 
			id, 
			user_id, 
			crypto, 
			type, 
			generated_at, 
			expires_at, 
			modified_at
		FROM security_tokens 
		WHERE id = $1;`

	row := r.session.QueryRow(q, id.String())
	// Fixed the Scan mapping
	if err := row.Scan(
		&i,
		&uid,
		&s.Crypto,
		&s.Type,
		&s.GeneratedAt,
		&s.ExpiresAt,
		&s.ModifiedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No token found
		}
		return nil, fmt.Errorf("error scanning SecurityToken by ID: %s", err)
	}

	// Parse ULIDs correctly
	s.Id = ulid.MustParse(i)
	s.UserId = ulid.MustParse(uid)

	return &s, nil
}

func (r *PostgresSecurityTokenRepository) GetByCrypto(crypto crypto.Crypto) (*security_token.SecurityToken, error) {
	var s security_token.SecurityToken
	var i string
	var uid string
	q := `SELECT 
			id, 
			user_id, 
			crypto, 
			type, 
			generated_at, 
			expires_at, 
			modified_at
		FROM security_tokens 
		WHERE crypto = $1;`

	row := r.session.QueryRow(q, crypto)
	if err := row.Scan(
		&i,
		&uid,
		&s.Crypto,
		&s.Type,
		&s.GeneratedAt,
		&s.ExpiresAt,
		&s.ModifiedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning SecurityToken by Crypto: %s", err)
	}

	s.Id = ulid.MustParse(i)
	s.UserId = ulid.MustParse(uid)

	return &s, nil
}

func (r *PostgresSecurityTokenRepository) GetAllForUser(user user.User) (*[]security_token.SecurityToken, error) {
	var securityTokens []security_token.SecurityToken
	q := `SELECT 
			id, 
			user_id, 
			crypto, 
			type, 
			generated_at, 
			expires_at, 
			modified_at
		FROM security_tokens 
		WHERE user_id = $1`

	rows, err := r.session.Query(q, user.Id.String())
	if err != nil {
		return nil, fmt.Errorf("error fetching all SecurityTokens for User: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var s security_token.SecurityToken
		var i string
		var uid string

		if err := rows.Scan(&i, &uid, &s.Crypto, &s.Type, &s.GeneratedAt, &s.ExpiresAt, &s.ModifiedAt); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, fmt.Errorf("error scanning all SecurityTokens for User: %s", err)
		}

		s.Id = ulid.MustParse(i)
		s.UserId = ulid.MustParse(uid)
		securityTokens = append(securityTokens, s)
	}

	return &securityTokens, nil
}

func (r *PostgresSecurityTokenRepository) ClearAllForUser(user user.User) error {
	q := `DELETE FROM security_tokens WHERE user_id = $1;`

	_, err := r.session.Exec(q, user.Id.String())
	if err != nil {
		return fmt.Errorf("error clearing all SecurityTokens for User: %s", err)
	}

	return nil
}

func (r *PostgresSecurityTokenRepository) Add(securityToken *security_token.SecurityToken) error {
	q := `INSERT INTO security_tokens (
			id, 
			user_id, 
			crypto, 
			type, 
			generated_at, 
			expires_at, 
			modified_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7);`

	_, err := r.session.Exec(
		q,
		securityToken.Id.String(),
		securityToken.UserId.String(),
		securityToken.Crypto,
		securityToken.Type,
		securityToken.GeneratedAt,
		securityToken.ExpiresAt,
		securityToken.ModifiedAt,
	)
	if err != nil {
		return fmt.Errorf("error adding SecurityToken: %s", err)
	}

	return nil
}

func (r *PostgresSecurityTokenRepository) Remove(securityToken *security_token.SecurityToken) error {
	q := `DELETE FROM security_tokens WHERE id = $1;`

	_, err := r.session.Exec(q, securityToken.Id.String())
	if err != nil {
		return fmt.Errorf("error removing SecurityToken: %s", err)
	}

	return nil
}

func NewPostgresSecurityTokenRepository(session *sql.DB) security_token.Repository {
	return &PostgresSecurityTokenRepository{
		session: session,
	}
}
