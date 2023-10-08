package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func db_connect() *sql.DB {
	db, _ := sql.Open("sqlite3", ":memory:")
	return db
}

// func migrate_schema(db *sql.DB) *SQLiteRepository {
// 	repo := InitSQLiteRepository(db)
// 	_ = repo.Migrate()
// 	return repo
// }

func TestInitSQLiteRepository(t *testing.T) {
	db := db_connect()
	defer db.Close()

	repo := InitSQLiteRepository(db)
	if repo.DB == nil {
		t.Error("db should not be nil")
	}
}

func TestSQLiteRepositoryMigrate(t *testing.T) {
	db := db_connect()
	defer db.Close()

	repo := InitSQLiteRepository(db)
	if repo.DB == nil {
		t.Error("db should not be nil")
	}
	err := repo.Migrate()
	if err != nil {
		t.Errorf("\nunable to migrate: %s\n", err.Error())
	}
	_, err = db.Exec("SELECT * FROM files")
	if err != nil {
		t.Errorf("\nunable to query table files: %s\n", err)
	}
	_, err = db.Exec("SELECT * FROM accounts")
	if err != nil {
		t.Errorf("\nunable to query table accounts: %s\n", err)
	}
}

func TestSQLiteRepositoryAccountCreate(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	if repo.DB == nil {
		t.Error("db should not be nil")
	}
	repo.Migrate()
	user, err := repo.CreateAccount("test@test.com", "test name", "password")
	if err != nil {
		t.Errorf("\nunable to create account: %s", err.Error())
	}
	if user.Email != "test@test.com" {
		t.Error("account should have correct email")
	}
	if user.Name != "test name" {
		t.Error("account should have correct name")
	}
	if user.PasswordHash != "password" {
		t.Error("account should have correct password hash")
	}
	if user.ID == 0 {
		t.Errorf("\naccount id should not be equal to 0: %d\n", user.ID)
	}
}

func TestSQLiteRepositoryGetAccount(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	if repo.DB == nil {
		t.Error("db should not be nil")
	}
	repo.Migrate()
	new_account, err := repo.CreateAccount("test@test.com", "test name", "password")
	if err != nil {
		t.Errorf("unable to create account: %s\n", err.Error())
	}

	account, err := repo.GetAccount(new_account.ID)
	if err != nil {
		t.Errorf("unable to get account: %s\n", err.Error())
	}

	if account.Email != "test@test.com" {
		t.Error("account should have \"test@test.com\" email")
	}

	if account.Name != "test name" {
		t.Error("account should have \"test name\" name")
	}

	if account.PasswordHash != "password" {
		t.Error("account should have \"password\" password")
	}
}

func TestSQLiteRepositoryChangeAccount(t *testing.T) {
	db := db_connect()
	defer db.Close()
	repo := InitSQLiteRepository(db)
	if repo.DB == nil {
		t.Error("db should not be nil")
	}
	repo.Migrate()
	account, _ := repo.CreateAccount("test@test.com", "test name", "password")

	account.Name = "new name"
	account.Email = "new@email.com"
	account.PasswordHash = "new password"

	err := repo.ChangeAccount(account.ID, account)
	if err != nil {
		t.Errorf("unable to change account: %s\n", err.Error())
	}

	updated_account, _ := repo.GetAccount(account.ID)
	if updated_account.Name != "new name" {
		t.Errorf("Account should has updated name \"new name\", but got: %s", updated_account.Name)
	}
	if updated_account.Email != "new@email.com" {
		t.Errorf("Account should has updated email \"new@email.com\", but got: %s", updated_account.Email)
	}
	if updated_account.PasswordHash != "new password" {
		t.Errorf("Account should has updated password \"new password\", but got: %s", updated_account.PasswordHash)
	}
}
