package database

type DatabaseRepository interface {
	AccountCreate(email, name, password string) (*User, error)
	AccountChange(user *User) error
	SaveFile(file *File) error
	GetFile(file_id string) (*File, error)
	DeleteFile(file_id string) error
	ChangeFile(file *File) error
	CreateFolder(folder *Folder) error
	ChangeFolder(folder *Folder) error
	DeleteFolder(folder_id string) error
}
