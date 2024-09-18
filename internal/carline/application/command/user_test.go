package command

import (
	"github.com/oklog/ulid/v2"
	"testing"
)

func TestThatCommandNameReturnsExpectedValueUpdateUserEmailAddress(t *testing.T) {
	cmd := UpdateUserEmailAddress{
		Id:           ulid.Make(),
		EmailAddress: "thomas@allen.com",
	}

	if cmd.CommandName() != "UpdateUserEmailAddress" {
		t.Fatal("test assertion failed for UpdateUserEmailAddress.CommandName()")
	}
}
