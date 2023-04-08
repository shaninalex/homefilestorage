package database

import (
	"database/sql"
	"errors"
	"homestorage/app/utils"
)

type Repository interface {
	Migrate() error
	GetUser(id int) (*utils.User, error)
	GetUserByEmail(email string) (utils.User, error)
	CreateUser(payload *utils.CreateUserPayload) error
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
