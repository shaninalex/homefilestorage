package typedefs

import "time"

type File struct {
	ID         uint      `json:"id,omitempty"`
	Name       string    `json:"name"`
	MimeType   string    `json:"mime_type"`
	Size       uint      `json:"size"`
	SystemPath string    `json:"system_path"`
	Owner      string    `json:"owner"`
	Hash       string    `json:"hash"`
	Public     bool      `json:"public"`
	FolderId   uint      `json:"folder_id"`
	Created_at time.Time `json:"created_at,omitempty"`
}
