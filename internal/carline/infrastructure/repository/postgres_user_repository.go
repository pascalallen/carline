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

func NewPostgresUserRepository(session *sql.DB) user.Repository {
	return &PostgresUserRepository{
		session: session,
	}
}

func (r *PostgresUserRepository) GetById(id ulid.ULID) (*user.User, error) {
	var u user.User
	var i string
	q := `SELECT 
			id,
			first_name,
			last_name,
			email_address,
			password_hash,
			created_at,
			modified_at
		FROM users
		WHERE id = $1`

	row := r.session.QueryRow(q, id.String())
	if err := row.Scan(&i, &u.FirstName, &u.LastName, &u.EmailAddress, &u.PasswordHash, &u.CreatedAt, &u.ModifiedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning User by ID: %s", err)
	}

	u.Id = ulid.MustParse(i)

	return &u, nil
}

func (r *PostgresUserRepository) GetByEmailAddress(emailAddress string) (*user.User, error) {
	var u user.User
	var id string
	q := `SELECT 
			id,
			first_name,
			last_name,
			email_address,
			password_hash,
			created_at,
			modified_at
		FROM users 
		WHERE email_address = $1`

	row := r.session.QueryRow(q, emailAddress)
	if err := row.Scan(&id, &u.FirstName, &u.LastName, &u.EmailAddress, &u.PasswordHash, &u.CreatedAt, &u.ModifiedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning User by email address: %s", err)
	}

	u.Id = ulid.MustParse(id)

	return &u, nil
}

func (r *PostgresUserRepository) GetAll(schoolId ulid.ULID) (*[]user.User, error) {
	var users []user.User
	q := `SELECT 
			id,
			first_name,
			last_name,
			email_address,
			password_hash,
			created_at,
			modified_at
		FROM users u
		JOIN user_schools us ON us.user_id = u.id
		WHERE us.school_id = $1`

	rows, err := r.session.Query(q, schoolId.String())
	if err != nil {
		return nil, fmt.Errorf("error fetching all Users: %s", err)
	}

	for rows.Next() {
		var id string
		var u user.User

		if err := rows.Scan(&id, &u.FirstName, &u.LastName, &u.EmailAddress, &u.PasswordHash, &u.CreatedAt, &u.ModifiedAt); err != nil {
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
	q := `INSERT INTO users(id, first_name, last_name, email_address, password_hash, created_at) VALUES($1, $2, $3, $4, $5, $6)`

	if _, err := r.session.Exec(q, user.Id.String(), user.FirstName, user.LastName, user.EmailAddress, user.PasswordHash, user.CreatedAt); err != nil {
		return fmt.Errorf("failed to persist User to database: %v", err)
	}

	if len(user.Roles) > 0 {
		q := `INSERT INTO user_roles(user_id, role_id) VALUES($1, $2)`
		for _, role := range user.Roles {
			if _, err := r.session.Exec(q, user.Id.String(), role.Id.String()); err != nil {
				return fmt.Errorf("failed to persist User Role to database: %v", err)
			}
		}
	}

	if len(user.Schools) > 0 {
		q := `INSERT INTO user_schools(user_id, school_id) VALUES($1, $2)`
		for _, school := range user.Schools {
			if _, err := r.session.Exec(q, user.Id.String(), school.Id.String()); err != nil {
				return fmt.Errorf("failed to persist User-School association to database: %v", err)
			}
		}
	}

	return nil
}

func (r *PostgresUserRepository) Remove(user *user.User) error {
	q := `DELETE FROM users WHERE id = $1`

	if _, err := r.session.Exec(q, user.Id); err != nil {
		return fmt.Errorf("failed to remove User from database: %v", err)
	}

	return nil
}

func (r *PostgresUserRepository) Save(user *user.User) error {
	q := `UPDATE users
		  SET first_name = $1,
			  last_name = $2,
			  email_address = $3,
			  password_hash = $4,
			  modified_at = $5
		  WHERE id = $6`

	res, err := r.session.Exec(
		q,
		user.FirstName,
		user.LastName,
		user.EmailAddress,
		user.PasswordHash,
		user.ModifiedAt,
		user.Id.String(),
	)

	if err != nil {
		return fmt.Errorf("failed to update User in database: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to verify update operation: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no User found with id: %s", user.Id.String())
	}

	return nil
}
