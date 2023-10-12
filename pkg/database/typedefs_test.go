package database

import "testing"

func TestAccountCheckPassword(t *testing.T) {

	account := &Account{
		ID:           1,
		Email:        "test@test.com",
		Name:         "test",
		PasswordHash: "password", // TODO: actual hash password
	}

	if !account.CheckPassword("password") {
		t.Error("Account.CheckPassword does not success")
	}
}

func TestAccountHashPassword(t *testing.T) {

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
