package database

import (
	"database/sql"
	"errors"
	"time"
)

type Repository interface {
	Migrate() error
	GetUser(id int) (*User, error)
	GetUserByEmail(email string) (User, error)
	CreateUser(payload *CreateUserPayload) error
}

type CreateUserPayload struct {
	Email          string
	HashedPassword string
}

type User struct {
	Id              int
	Email           string
	Hashed_password string
	Active          bool
	Created_at      *time.Time
}

type DBConfig struct {
	DB_ENGINE string
	DB_PATH   string
}

type DatabaseRepository struct {
	db *sql.DB
}

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("record not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)
