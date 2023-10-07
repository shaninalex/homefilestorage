package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/shaninalex/homefilestorage/pkg/typedefs"
)

type Database struct {
	DB *sql.DB
}

func CreateConnection(db_path string) (*Database, error) {
	var database Database
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		return nil, err
	}
	database.DB = db
	return &database, nil
}

func (db *Database) FileSave(f *typedefs.File) error {
	err := db.DB.QueryRow(`
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

func (db *Database) GetFile(user_id string, file_id int64) (*typedefs.File, error) {
	var file typedefs.File
	err := db.DB.QueryRow(`SELECT * FROM files WHERE id = $1 AND user_id = $2`,
		file_id, user_id).Scan(
		&file.ID, &file.Name, &file.MimeType, &file.Size, &file.SystemPath,
		&file.Owner, &file.Hash, &file.Public, &file.FolderId, &file.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (db *Database) FileDelete(id string) (*typedefs.File, error) {
	// delete from database
	// rabbitmq request to storage to delete file
	return nil, nil
}

func (db *Database) GetUserFiles(user_id string, folder_id int64) ([]typedefs.File, error) {
	rows, err := db.DB.Query(`SELECT * FROM files WHERE user_id = $1 AND folder = $2`, user_id, folder_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []typedefs.File
	for rows.Next() {
		var file typedefs.File
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
