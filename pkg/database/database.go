package database

import (
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/mattn/go-sqlite3"
)

type Repository interface {
	Migrate() error
	GetAccount(id int64) (*Account, error)
	GetAccountByEmail(email string) (*Account, error)
	CreateAccount(email, name, password string) (*Account, error)
	ChangeAccount(id int64, user *Account) error
	SaveFile(file *File) error
	GetFile(file_id int64) (*File, error)
	DeleteFile(file_id int64) error
	ChangeFile(file *File) error
	GetFolder(folder_id int64) (*Folder, error)
	CreateFolder(folder *Folder) error
	ChangeFolder(folder *Folder) error
	DeleteFolder(folder_id int64) error
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
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		CHECK ("size" > 0)
	);
	CREATE TABLE IF NOT EXISTS accounts
	(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		password_hash TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
	);
	CREATE TABLE IF NOT EXISTS folders
	(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		color VARCHAR(10),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
	)
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

func (db *SQLiteRepository) GetAccountByEmail(email string) (*Account, error) {
	selectSQL, _, _ := goqu.From("accounts").Select(
		"id", "name", "email", "password_hash",
	).Where(
		goqu.C("email").Eq(email),
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

	if account.Name == "" {
		// TODO: tmp
		return nil, errors.New("user not found")
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
	insertSQL, _, _ := goqu.Insert("files").Rows(
		goqu.Record{
			"name":        file.Name,
			"mime_type":   file.MimeType,
			"size":        file.Size,
			"system_path": file.SystemPath,
			"folder_id":   file.FolderId,
			"created_at":  file.CreatedAt,
		},
	).Returning(goqu.C("id")).ToSQL()
	err := db.DB.QueryRow(insertSQL).Scan(&file.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteRepository) GetFile(file_id int64) (*File, error) {
	selectSQL, _, _ := goqu.From("files").Select(
		"id", "name", "mime_type", "size", "system_path", "folder_id", "created_at",
	).Where(
		goqu.C("id").Eq(file_id),
	).ToSQL()
	var file File
	err := db.DB.QueryRow(selectSQL).Scan(
		&file.ID,
		&file.Name,
		&file.MimeType,
		&file.Size,
		&file.SystemPath,
		&file.FolderId,
		&file.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (db *SQLiteRepository) DeleteFile(file_id int64) error {
	deleteSQL, _, _ := goqu.Delete("files").Where(goqu.C("id").Eq(file_id)).ToSQL()
	_, err := db.DB.Exec(deleteSQL)
	if err != nil {
		return err
	}
	return nil
}

// "move" file from folder or IN folder is the same as ChangeFile. So it's
// unnececery to create separate handlers
func (db *SQLiteRepository) ChangeFile(file *File) error {
	updateSQL, _, _ := goqu.Update("files").Set(
		goqu.Record{
			"name":      file.Name,
			"folder_id": file.FolderId,
		},
	).Where(
		goqu.C("id").Eq(file.ID),
	).ToSQL()
	_, err := db.DB.Exec(updateSQL)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteRepository) GetFolder(folder_id int64) (*Folder, error) {
	selectSQL, _, _ := goqu.From("folders").Select(
		"id", "name", "color",
	).Where(goqu.C("id").Eq(folder_id)).ToSQL()
	var folder Folder
	err := db.DB.QueryRow(selectSQL).Scan(
		&folder.ID,
		&folder.Name,
		&folder.Color,
	)
	if err != nil {
		return nil, err
	}
	return &folder, nil
}

func (db *SQLiteRepository) CreateFolder(folder *Folder) error {
	insertSQL, _, _ := goqu.Insert("folders").Rows(
		goqu.Record{
			"name":  folder.Name,
			"color": folder.Color,
		},
	).Returning(goqu.C("id")).ToSQL()
	err := db.DB.QueryRow(insertSQL).Scan(&folder.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteRepository) ChangeFolder(folder *Folder) error {
	updateSQL, _, _ := goqu.Update("folders").Set(
		goqu.Record{
			"name":  folder.Name,
			"color": folder.Color,
		},
	).Where(
		goqu.C("id").Eq(folder.ID),
	).ToSQL()
	_, err := db.DB.Exec(updateSQL)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteRepository) DeleteFolder(folder_id int64) error {
	deleteSQL, _, _ := goqu.Delete("folders").Where(goqu.C("id").Eq(folder_id)).ToSQL()
	_, err := db.DB.Exec(deleteSQL)
	if err != nil {
		return err
	}
	return nil

}
