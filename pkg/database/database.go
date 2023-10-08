package database

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type Repository interface {
	Migrate() error
	GetAccount(id int64) (*Account, error)
	CreateAccount(email, name, password string) (*Account, error)
	ChangeAccount(id int64, user *Account) error
	SaveFile(file *File) error
	GetFile(file_id string) (*File, error)
	DeleteFile(file_id string) error
	ChangeFile(file *File) error
	CreateFolder(folder *Folder) error
	ChangeFolder(folder *Folder) error
	DeleteFolder(folder_id string) error
}

type SQLiteRepository struct {
	DB *sql.DB
}

func InitSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		DB: db,
	}
}

func (db *SQLiteRepository) Migrate() error {
	scheme := `
	CREATE TABLE IF NOT EXISTS files
	(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		mime_type TEXT NOT NULL,
		size INTEGER NOT NULL,
		system_path TEXT,
		folder_id INTEGER,
		created_at TIMESTAMP,
		CHECK ("size" > 0)
	);
	CREATE TABLE IF NOT EXISTS accounts
	(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		password_hash TEXT NOT NULL
	);
	`
	_, err := db.DB.Exec(scheme)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteRepository) GetAccount(id int64) (*Account, error) {
	selectSQL, _, _ := goqu.From("accounts").Select(
		"id", "name", "email", "password_hash",
	).Where(
		goqu.C("id").Eq(id),
	).ToSQL()
	var account Account
	err := db.DB.QueryRow(selectSQL).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.PasswordHash,
	)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (db *SQLiteRepository) CreateAccount(email, name, password string) (*Account, error) {
	account := Account{Email: email, Name: name, PasswordHash: password}
	insertSQL, _, _ := goqu.Insert("accounts").Rows(
		goqu.Record{
			"email":         account.Email,
			"name":          account.Name,
			"password_hash": account.PasswordHash,
		},
	).Returning(goqu.C("id")).ToSQL()
	err := db.DB.QueryRow(insertSQL).Scan(&account.ID)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (db *SQLiteRepository) ChangeAccount(id int64, user *Account) error {
	updateAccountSQL, _, _ := goqu.Update("accounts").Set(
		goqu.Record{
			"email":         user.Email,
			"name":          user.Name,
			"password_hash": user.PasswordHash,
		},
	).Where(
		goqu.C("id").Eq(id),
	).ToSQL()
	_, err := db.DB.Exec(updateAccountSQL)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteRepository) SaveFile(file *File) error {
	return nil
}

func (db *SQLiteRepository) GetFile(file_id string) (*File, error) {
	return nil, nil
}

func (db *SQLiteRepository) DeleteFile(file_id string) error {
	return nil
}

func (db *SQLiteRepository) ChangeFile(file *File) error {
	return nil
}

func (db *SQLiteRepository) CreateFolder(folder *Folder) error {
	return nil
}

func (db *SQLiteRepository) ChangeFolder(folder *Folder) error {
	return nil
}

func (db *SQLiteRepository) DeleteFolder(folder_id string) error {
	return nil
}
