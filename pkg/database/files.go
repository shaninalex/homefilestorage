package database

import "time"

type File struct {
	ID         int64     `json:"id,omitempty" db:"id"`
	Name       string    `json:"name" db:"name"`
	MimeType   string    `json:"mime_type" db:"mime_type"`
	Size       int64     `json:"size" db:"size"`
	SystemPath string    `json:"system_path" db:"system_path"`
	FolderId   int64     `json:"folder_id" db:"folder_id"`
	Created_at time.Time `json:"created_at,omitempty" db:"created_at"`
}

func (f *File) GetFileSize() int64 {
	return 0
}

func (f *File) GetMimeType() int64 {
	return 0
}

type Folder struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}

func (f *Folder) GetFolderSize() int64 {
	return 0
}
