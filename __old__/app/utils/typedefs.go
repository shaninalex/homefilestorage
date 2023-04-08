package utils

import (
	"time"
)

type CreateUserRequestPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type LoginUserRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshToken struct {
	Access string `json:"access"`
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

type BooleanResponse struct {
	Status bool `json:"status"`
}

type File struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	MimeType   string    `json:"mime_type"`
	Size       int       `json:"size"`
	SystemPath string    `json:"system_path"`
	Owner      int       `json:"owner"`
	Hash       string    `json:"hash"`
	Public     bool      `json:"public"`
	Folder     *int      `json:"folder"`
	Created_at time.Time `json:"created_at"`
}

type Folder struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Color      string    `json:"color"`
	Owner      int       `json:"owner"`
	Created_at time.Time `json:"created_at"`
}

type CreateUserPayload struct {
	Email          string
	HashedPassword string
}

type User struct {
	Id              int        `json:"id"`
	Email           string     `json:"email"`
	Hashed_password string     `json:"password"`
	Active          bool       `json:"active"`
	Created_at      *time.Time `json:"created_at"`
}

type DBConfig struct {
	DB_ENGINE string
	DB_PATH   string
}
