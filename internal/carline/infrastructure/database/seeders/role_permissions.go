package seeders

import (
	"database/sql"
	"fmt"
)

var rolePermissions = []struct {
	RoleName       string
	PermissionName string
}{
	// SUPER ADMIN: Full access to everything
	{"ROLE_SUPER_ADMIN", "CREATE_ROLES"},
	{"ROLE_SUPER_ADMIN", "READ_ROLES"},
	{"ROLE_SUPER_ADMIN", "UPDATE_ROLES"},
	{"ROLE_SUPER_ADMIN", "DELETE_ROLES"},
	{"ROLE_SUPER_ADMIN", "MANAGE_ROLE_PERMISSIONS"},
	{"ROLE_SUPER_ADMIN", "CREATE_USERS"},
	{"ROLE_SUPER_ADMIN", "READ_USERS"},
	{"ROLE_SUPER_ADMIN", "UPDATE_USERS"},
	{"ROLE_SUPER_ADMIN", "DELETE_USERS"},
	{"ROLE_SUPER_ADMIN", "MANAGE_USER_ROLES"},
	{"ROLE_SUPER_ADMIN", "CREATE_SCHOOLS"},
	{"ROLE_SUPER_ADMIN", "READ_SCHOOLS"},
	{"ROLE_SUPER_ADMIN", "UPDATE_SCHOOLS"},
	{"ROLE_SUPER_ADMIN", "DELETE_SCHOOLS"},
	{"ROLE_SUPER_ADMIN", "CREATE_STUDENTS"},
	{"ROLE_SUPER_ADMIN", "READ_STUDENTS"},
	{"ROLE_SUPER_ADMIN", "UPDATE_STUDENTS"},
	{"ROLE_SUPER_ADMIN", "DELETE_STUDENTS"},
	{"ROLE_SUPER_ADMIN", "IMPORT_STUDENTS"},
	{"ROLE_SUPER_ADMIN", "READ_PERMISSIONS"},
	{"ROLE_SUPER_ADMIN", "UPDATE_PERMISSIONS"},

	// ADMIN: CRUD for schools, its users, and its students
	{"ROLE_ADMIN", "CREATE_SCHOOLS"},
	{"ROLE_ADMIN", "READ_SCHOOLS"},
	{"ROLE_ADMIN", "UPDATE_SCHOOLS"},
	{"ROLE_ADMIN", "DELETE_SCHOOLS"},
	{"ROLE_ADMIN", "CREATE_USERS"},
	{"ROLE_ADMIN", "READ_USERS"},
	{"ROLE_ADMIN", "UPDATE_USERS"},
	{"ROLE_ADMIN", "DELETE_USERS"},
	{"ROLE_ADMIN", "CREATE_STUDENTS"},
	{"ROLE_ADMIN", "READ_STUDENTS"},
	{"ROLE_ADMIN", "UPDATE_STUDENTS"},
	{"ROLE_ADMIN", "DELETE_STUDENTS"},
	{"ROLE_ADMIN", "IMPORT_STUDENTS"},

	// USER: CRUD for their school's students
	{"ROLE_USER", "CREATE_STUDENTS"},
	{"ROLE_USER", "READ_STUDENTS"},
	{"ROLE_USER", "UPDATE_STUDENTS"},
	{"ROLE_USER", "DELETE_STUDENTS"},
}

func SeedRolePermissions(db *sql.DB) error {
	for _, rp := range rolePermissions {
		var roleId string
		roleQuery := `SELECT id FROM roles WHERE name = $1`
		if err := db.QueryRow(roleQuery, rp.RoleName).Scan(&roleId); err != nil {
			return fmt.Errorf("failed to fetch role ID for %s: %v", rp.RoleName, err)
		}

		var permissionId string
		permissionQuery := `SELECT id FROM permissions WHERE name = $1`
		if err := db.QueryRow(permissionQuery, rp.PermissionName).Scan(&permissionId); err != nil {
			return fmt.Errorf("failed to fetch permission ID for %s: %v", rp.PermissionName, err)
		}

		insertQuery := `INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2) ON CONFLICT (role_id, permission_id) DO NOTHING`
		if _, err := db.Exec(insertQuery, roleId, permissionId); err != nil {
			return fmt.Errorf("failed to seed role_permissions: %v", err)
		}
	}

	return nil
}
