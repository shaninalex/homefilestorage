package database

import "time"

type File struct {
	ID         uint      `json:"id,omitempty"`
	Name       string    `json:"name"`
	MimeType   string    `json:"mime_type"`
	Size       uint      `json:"size"`
	SystemPath string    `json:"system_path"`
	Owner      uint      `json:"owner"` // foreign key to users table
	Hash       string    `json:"hash"`
	Public     bool      `json:"public"`
	FolderId   Folder    `json:"folder_id"`
	Created_at time.Time `json:"created_at,omitempty"`
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
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Owner     uint      `json:"owner"`
	CreatedAt time.Time `json:"created_at"`
}
