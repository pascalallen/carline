package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/permission"
)

type PostgresPermissionRepository struct {
	session *sql.DB
}

func NewPostgresPermissionRepository(session *sql.DB) permission.Repository {
	return &PostgresPermissionRepository{
		session: session,
	}
}

func (r *PostgresPermissionRepository) GetById(id ulid.ULID) (*permission.Permission, error) {
	var p permission.Permission
	var i string

	row := r.session.QueryRow("SELECT * FROM permissions WHERE id = $1", id.String())
	if err := row.Scan(&i, &p.Name, &p.Description, &p.CreatedAt, &p.ModifiedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning Permission by ID: %s", err)
	}

	p.Id = ulid.MustParse(i)

	return &p, nil
}

func (r *PostgresPermissionRepository) GetByName(name string) (*permission.Permission, error) {
	var p permission.Permission
	var i string

	row := r.session.QueryRow("SELECT * FROM permissions WHERE name = $1", name)
	if err := row.Scan(&i, &p.Name, &p.Description, &p.CreatedAt, &p.ModifiedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning Permission by name: %s", err)
	}

	p.Id = ulid.MustParse(i)

	return &p, nil
}

func (r *PostgresPermissionRepository) GetAll() (*[]permission.Permission, error) {
	var p permission.Permission
	var permissions []permission.Permission
	var id string

	rows, err := r.session.Query("SELECT * FROM permissions")
	if err != nil {
		return nil, fmt.Errorf("error fetching all Permissions: %s", err)
	}

	for rows.Next() {
		if err := rows.Scan(&id, &p.Name, &p.Description, &p.CreatedAt, &p.ModifiedAt); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, fmt.Errorf("error scanning all Permissions: %s", err)
		}

		p.Id = ulid.MustParse(id)
		permissions = append(permissions, p)
	}

	return &permissions, nil
}

func (r *PostgresPermissionRepository) Add(permission *permission.Permission) error {
	if _, err := r.session.Exec("INSERT INTO permissions(id, name, description, created_at) VALUES($1, $2, $3, $4)", permission.Id.String(), permission.Name, permission.Description, permission.CreatedAt); err != nil {
		return fmt.Errorf("failed to persist Permission to database: %v", err)
	}

	return nil
}

func (r *PostgresPermissionRepository) Remove(permission *permission.Permission) error {
	if _, err := r.session.Exec("DELETE FROM permissions WHERE id = $1", permission.Id.String()); err != nil {
		return fmt.Errorf("failed to delete Permission from database: %s", permission)
	}

	return nil
}
