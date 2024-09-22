package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/user"
)

type PostgresUserRepository struct {
	session *sql.DB
}

func NewGormUserRepository(session *sql.DB) user.Repository {
	return &PostgresUserRepository{
		session: session,
	}
}

func (r *PostgresUserRepository) GetById(id ulid.ULID) (*user.User, error) {
	var u user.User
	var i string

	row := r.session.QueryRow("SELECT * FROM users WHERE id = $1", id.String())
	if err := row.Scan(&i, &u.FirstName, &u.LastName, &u.EmailAddress, &u.PasswordHash, &u.CreatedAt, &u.ModifiedAt, &u.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning User by ID: %s", err)
	}

	u.Id = ulid.MustParse(i)

	// TODO: eager load schools and roles

	return &u, nil
}

func (r *PostgresUserRepository) GetByEmailAddress(emailAddress string) (*user.User, error) {
	var u user.User
	var id string

	row := r.session.QueryRow("SELECT * FROM users WHERE email_address = $1", emailAddress)
	if err := row.Scan(&id, &u.FirstName, &u.LastName, &u.EmailAddress, &u.PasswordHash, &u.CreatedAt, &u.ModifiedAt, &u.DeletedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning User by email address: %s", err)
	}

	u.Id = ulid.MustParse(id)

	// TODO: eager load schools and roles

	return &u, nil
}

// GetAll TODO: Add pagination
func (r *PostgresUserRepository) GetAll(includeDeleted bool) (*[]user.User, error) {
	var u user.User
	var users []user.User
	var id string
	query := "SELECT * FROM users"

	if !includeDeleted {
		query += " WHERE deleted_at IS NULL"
	}

	rows, err := r.session.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all Users: %s", err)
	}

	for rows.Next() {
		if err := rows.Scan(&id, &u.FirstName, &u.LastName, &u.EmailAddress, &u.PasswordHash, &u.CreatedAt, &u.ModifiedAt, &u.DeletedAt); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, fmt.Errorf("error scanning all Users: %s", err)
		}

		u.Id = ulid.MustParse(id)
		users = append(users, u)
	}

	return &users, nil
}

func (r *PostgresUserRepository) Add(user *user.User) error {
	if _, err := r.session.Exec("INSERT INTO users(id, first_name, last_name, email_address, password_hash, created_at) VALUES($1, $2, $3, $4, $5, $6)", user.Id.String(), user.FirstName, user.LastName, user.EmailAddress, user.PasswordHash, user.CreatedAt); err != nil {
		return fmt.Errorf("failed to persist User to database: %v", err)
	}

	return nil
}

func (r *PostgresUserRepository) Remove(user *user.User) error {
	user.Delete()

	if _, err := r.session.Exec("UPDATE users SET deleted_at = $1 WHERE id = $2", user.DeletedAt, user.Id); err != nil {
		return fmt.Errorf("failed to soft delete User: %v", err)
	}

	return nil
}
