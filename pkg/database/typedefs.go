package database

import (
	"errors"
	"time"
)

type File struct {
	ID         int64     `json:"id,omitempty" db:"id"`
	Name       string    `json:"name" db:"name"`
	MimeType   string    `json:"mime_type" db:"mime_type"`
	Size       int64     `json:"size" db:"size"`
	SystemPath string    `json:"system_path" db:"system_path"`
	FolderId   int64     `json:"folder_id" db:"folder_id"`
	CreatedAt  time.Time `json:"created_at,omitempty" db:"created_at"`
}

func (f *File) GetFileSize() int64 {
	return 0
}

func (f *File) GetMimeType() int64 {
	return 0
}

func (f *File) FormatTime() string {
	return f.CreatedAt.Format("2006.01.02 15:04:05")
}

type Folder struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Color     string    `json:"color" db:"color"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Account struct {
	ID           int64  `json:"-" db:"id"`
	Email        string `json:"email" db:"email"`
	Name         string `json:"name" db:"name"`
	PasswordHash string `json:"-" db:"password_hash"`
}

func (a *Account) CheckPassword(clear_password string) bool {
	return false
}

func (a *Account) HashPassword(clear_password string) error {
	return errors.New("not implemented")
}
