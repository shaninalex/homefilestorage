package database

import (
	"fmt"
	"log"
	"time"

	"github.com/matthewhartstonge/argon2"
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

func (f *File) FormatTime() string {
	return f.CreatedAt.Format("2006.01.02 15:04:05")
}

func (f *File) FormatSize() string {
	if f.Size < 1024 {
		return fmt.Sprintf("%.1f Kb", float64(f.Size)/1024)
	} else if f.Size < 1024*1024*1024 {
		return fmt.Sprintf("%.1f Mb", float64(f.Size)/(1024*1024))
	} else {
		return fmt.Sprintf("%.1f Gb", float64(f.Size)/(1024*1024*1024))
	}
}

func (f *File) DownloadPath() string {
	return fmt.Sprintf("/files/%d", f.ID)
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
	ok, err := argon2.VerifyEncoded([]byte(clear_password), []byte(a.PasswordHash))
	if err != nil {
		log.Println(err)
		return false
	}
	if ok {
		return true
	}
	return false
}

func (a *Account) HashPassword(clear_password string) error {
	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(clear_password))
	if err != nil {
		return err
	}
	a.PasswordHash = string(encoded)
	return nil
}
