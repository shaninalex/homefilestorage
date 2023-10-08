package database

import "time"

type File struct {
	ID         uint      `json:"id,omitempty"`
	Name       string    `json:"name"`
	MimeType   string    `json:"mime_type"`
	Size       uint      `json:"size"`
	SystemPath string    `json:"system_path"`
	FolderId   uint      `json:"folder_id"`
	Created_at time.Time `json:"created_at,omitempty"`
}

func (f *File) GetFileSize() int64 {
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
