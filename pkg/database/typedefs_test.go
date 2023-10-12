package database

import "testing"

func TestAccountPassword(t *testing.T) {

	account := &Account{
		ID:    1,
		Email: "test@test.com",
		Name:  "test",
	}

	err := account.HashPassword("password")
	if err != nil {
		t.Errorf("Account.HashPassword return error: %s\n", err)
	}

	if !account.CheckPassword("password") {
		t.Error("Account.CheckPassword does not success")
	}
}
