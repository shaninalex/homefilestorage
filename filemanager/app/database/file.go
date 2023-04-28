package database

import "time"

type File struct {
	ID         uint      `json:"id,omitempty" gorm:"id"`
	Name       string    `json:"name" gorm:"name"`
	MimeType   string    `json:"mime_type" gorm:"mime_type"`
	Size       uint      `json:"size" gorm:"size"`
	SystemPath string    `json:"system_path" gorm:"system_path"`
	Owner      uint      `json:"owner" gorm:"owner"` // foreign key to users table
	Hash       string    `json:"hash" gorm:"hash"`
	Public     bool      `json:"public" gorm:"public"`
	FolderId   Folder    `json:"folder_id" gorm:"references:ID"`
	Created_at time.Time `json:"created_at,omitempty" gorm:"created_at"`
}

func FileSave(f *File) (*File, error) {
	return nil, nil
}

func FileGet(id string) (*File, error) {
	return nil, nil
}

func FileDelete(id string) (*File, error) {
	// rabbitmq request to storage to delete file
	return nil, nil
}

func GetUserFiles(user_id string, folder_id *uint) ([]File, error) {
	// If folder_id == nil -> it's root directory.
	return nil, nil
}

type Folder struct {
	ID        uint      `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Color     string    `json:"color" gorm:"color"`
	Owner     uint      `json:"owner" gorm:"owner"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}
