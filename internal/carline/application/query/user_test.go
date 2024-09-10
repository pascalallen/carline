package query

import (
	"github.com/oklog/ulid/v2"
	"testing"
)

func TestThatQueryNameReturnsExpectedValueGetUserById(t *testing.T) {
	qry := GetUserById{
		Id: ulid.Make(),
	}

	if qry.QueryName() != "GetUserById" {
		t.Fatal("test assertion failed for GetUserById.QueryName()")
	}
}

func TestThatQueryNameReturnsExpectedValueGetUserByEmailAddress(t *testing.T) {
	qry := GetUserByEmailAddress{
		EmailAddress: "foo@bar.com",
	}

	if qry.QueryName() != "GetUserByEmailAddress" {
		t.Fatal("test assertion failed for GetUserByEmailAddress.QueryName()")
	}
}
