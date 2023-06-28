package database

import (
	"database/sql"
	"log"
	"time"
)

type Database struct {
	db *sql.DB
}

func CreateConnection(connection_string string) (*Database, error) {
	var database Database
	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		return nil, err
	}
	database.db = db
	return &database, nil
}

type File struct {
	ID         uint      `json:"id,omitempty"`
	Name       string    `json:"name"`
	MimeType   string    `json:"mime_type"`
	Size       uint      `json:"size"`
	SystemPath string    `json:"system_path"`
	Owner      uint      `json:"owner"`
	Hash       string    `json:"hash"`
	Public     bool      `json:"public"`
	FolderId   uint      `json:"folder_id"`
	Created_at time.Time `json:"created_at,omitempty"`
}

func (db *Database) FileSave(f *File) error {
	err := db.db.QueryRow(`
		INSERT INTO files (name, mime_type, size, system_path, user_id, hash, public, folder) 
		VALUES ( $1, $2, $3, $4, $5, $6, $7, $8 ) RETURNING id, created_at`,
		f.Name, f.MimeType, f.Size, f.SystemPath, f.Owner, f.Hash, f.Public, f.FolderId,
	).Scan(&f.ID, &f.Created_at)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (db *Database) GetFile(user_id string, file_id int64) (*File, error) {
	var file File
	err := db.db.QueryRow(`SELECT * FROM files WHERE id = $1 AND user_id = $2`,
		file_id, user_id).Scan(
		&file.ID, &file.Name, &file.MimeType, &file.Size, &file.SystemPath,
		&file.Owner, &file.Hash, &file.Public, &file.FolderId, &file.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (db *Database) FileDelete(id string) (*File, error) {
	// delete from database
	// rabbitmq request to storage to delete file
	return nil, nil
}

func (db *Database) GetUserFiles(user_id string, folder_id int64) ([]File, error) {
	log.Printf("Get files list for user %s\n", user_id)
	rows, err := db.db.Query(`SELECT * FROM files WHERE user_id = $1 AND folder = $2`, user_id, folder_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []File
	for rows.Next() {
		var file File
		if err := rows.Scan(
			&file.ID, &file.Name, &file.MimeType, &file.Size, &file.SystemPath,
			&file.Owner, &file.Hash, &file.Public, &file.FolderId, &file.Created_at,
		); err != nil {
			return files, err
		}
		files = append(files, file)
	}
	if err = rows.Err(); err != nil {
		return files, err
	}
	return files, nil
}

type Folder struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Owner     uint      `json:"owner"`
	CreatedAt time.Time `json:"created_at"`
}
