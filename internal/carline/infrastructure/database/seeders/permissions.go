package seeders

import (
	"database/sql"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/permission"
	"time"
)

var permissions = []permission.Permission{
	{Id: ulid.MustParse("01FY7XRMMKB4FA7G0Q0D9S8CDN"), Name: "CREATE_ROLES", Description: "Allows the user to create roles"},
	{Id: ulid.MustParse("01FY7XP5V2EPJZFG361WRHJDVK"), Name: "READ_ROLES", Description: "Allows the user to have read access to roles"},
	{Id: ulid.MustParse("01FY7XTB323SXWWJ757AY5QJ7H"), Name: "UPDATE_ROLES", Description: "Allows the user to update roles"},
	{Id: ulid.MustParse("01FY7XVSJQHAC040RMMA37ZTNR"), Name: "DELETE_ROLES", Description: "Allows the user to delete roles"},
	{Id: ulid.MustParse("01FY7XXQMW888MCBXH67HADFY4"), Name: "MANAGE_ROLE_PERMISSIONS", Description: "Allows the user to manage role permissions"},

	// Permissions for USERS
	{Id: ulid.MustParse("01FY7XRW3JSY2Y4Q8XRVDYSCZK"), Name: "CREATE_USERS", Description: "Allows the user to create users"},
	{Id: ulid.MustParse("01FY7XMMX83NKP6Y0BSDEJ1HQP"), Name: "READ_USERS", Description: "Allows the user to have read access to users"},
	{Id: ulid.MustParse("01FY7XTMCG9EWGWW0K2DBF4BJJ"), Name: "UPDATE_USERS", Description: "Allows the user to update users"},
	{Id: ulid.MustParse("01FY7XW2NRY5FKSKTQ748TAY0D"), Name: "DELETE_USERS", Description: "Allows the user to delete users"},
	{Id: ulid.MustParse("01FY7XYK0ZEDJ9Z4RBXZQD5FW4"), Name: "MANAGE_USER_ROLES", Description: "Allows the user to manage user roles"},

	// Permissions for SCHOOLS
	{Id: ulid.MustParse("01FY8BX0XR1MWF9W26VPCXYB7E"), Name: "CREATE_SCHOOLS", Description: "Allows the user to create schools"},
	{Id: ulid.MustParse("01FY8BX4ZYDJ6DW4N3V52M41Z0"), Name: "READ_SCHOOLS", Description: "Allows the user to read school data"},
	{Id: ulid.MustParse("01FY8BXAE3FVJ8BV8S3MXA58QV"), Name: "UPDATE_SCHOOLS", Description: "Allows the user to update school data"},
	{Id: ulid.MustParse("01FY8BXD6ARMZEWPQN69WMP8WN"), Name: "DELETE_SCHOOLS", Description: "Allows the user to delete schools"},

	// Permissions for STUDENTS
	{Id: ulid.MustParse("01FY8BXN0P0MG28GRG7AKPBB92"), Name: "CREATE_STUDENTS", Description: "Allows the user to create students"},
	{Id: ulid.MustParse("01FY8BXP78J8TTV1MV9XN8GAVH"), Name: "READ_STUDENTS", Description: "Allows the user to read student data"},
	{Id: ulid.MustParse("01FY8BXRZNT536R8JDM97E95Y4"), Name: "UPDATE_STUDENTS", Description: "Allows the user to update student data"},
	{Id: ulid.MustParse("01FY8BXTPHAJPKWW2H05MCXJQZ"), Name: "DELETE_STUDENTS", Description: "Allows the user to delete students"},
	{Id: ulid.MustParse("01FY8BY040SQ72NFNCE5ZVS3KE"), Name: "IMPORT_STUDENTS", Description: "Allows the user to import student data"},

	// General Permissions
	{Id: ulid.MustParse("01FY7XQ0F9XK4B74QR4BGTXNFD"), Name: "READ_PERMISSIONS", Description: "Allows the user to have read access to permissions"},
	{Id: ulid.MustParse("01FY7XT3SRMCD2RNB9X8WR7VYJ"), Name: "UPDATE_PERMISSIONS", Description: "Allows the user to update permissions"},
}

func SeedPermissions(db *sql.DB) error {
	for _, p := range permissions {
		q := `INSERT INTO permissions (id, name, description, created_at) VALUES ($1, $2, $3, $4) ON CONFLICT (name) DO NOTHING`
		if _, err := db.Exec(q, p.Id.String(), p.Name, p.Description, time.Now()); err != nil {
			return fmt.Errorf("failed to seed permissions: %v", err)
		}
	}
	return nil
}
