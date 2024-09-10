package event

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/password"
	"testing"
)

func TestThatEventNameReturnsExpectedValueUserRegistered(t *testing.T) {
	u := UserRegistered{
		Id:           ulid.Make(),
		FirstName:    "Pascal",
		LastName:     "Allen",
		EmailAddress: "pascal@allen.com",
		PasswordHash: password.Create("pa$$w0rd"),
	}

	if u.EventName() != "UserRegistered" {
		t.Fatal("test assertion failed for UserRegistered.EventName()")
	}
}

func TestThatEventNameReturnsExpectedValueUserEmailAddressUpdated(t *testing.T) {
	u := UserEmailAddressUpdated{
		Id:           ulid.Make(),
		EmailAddress: "pascal@allen.com",
	}

	if u.EventName() != "UserEmailAddressUpdated" {
		t.Fatal("test assertion failed for UserEmailAddressUpdated.EventName()")
	}
}
