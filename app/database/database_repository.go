package database

import (
	"database/sql"
	"homestorage/app/utils"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// TODO
// 	- udpate
// 	- delete
//  - return predefined errors

func CreateDatabaseConnection(conf *utils.DBConfig) (*DatabaseRepository, error) {

	db, err := sql.Open(conf.DB_ENGINE, conf.DB_PATH)
	if err != nil {
		log.Fatalf("Cant connect to database: %s", err)
		return nil, err
	}

	log.Printf("Connected to db: %s %s", conf.DB_ENGINE, conf.DB_PATH)

	return &DatabaseRepository{
		db: db,
	}, nil
}

func (r *DatabaseRepository) Migrate() error {
	file, err := os.Open("app/database/scheme/scheme.sql")
	if err != nil {
		log.Fatal("Cant open scheme.sql")
		return err
	}

	query, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Cant read file: %s", err)
		return err
	}

	if _, err := r.db.Exec(string(query)); err != nil {
		log.Fatalf("Error while migrations: %s", err)
		return err
	}

	return nil
}

func (r *DatabaseRepository) GetUser(id *int) (*utils.User, error) {
	var user utils.User
	row := r.db.QueryRow("SELECT id, email, active, hashed_password, created_at FROM users WHERE id=?;", &id)
	err := row.Scan(&user.Id, &user.Email, &user.Active, &user.Hashed_password, &user.Created_at)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *DatabaseRepository) GetUserByEmail(email string) (*utils.User, error) {
	var user utils.User
	row := r.db.QueryRow("SELECT id, email, active, hashed_password, created_at FROM users WHERE email=?;", email)
	err := row.Scan(&user.Id, &user.Email, &user.Active, &user.Hashed_password, &user.Created_at)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *DatabaseRepository) CreateUser(payload *utils.CreateUserPayload) error {
	// TODO: Remove generating password hash from database package!
	query := "INSERT INTO users (email, hashed_password) VALUES (?, ?) RETURNING id;"
	_, err := r.db.Exec(query, payload.Email, payload.HashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func (r *DatabaseRepository) SaveFileRecord(payload *utils.File) (int, error) {
	var id int64
	query := `
		INSERT INTO files 
			(name, mime_type, size, system_path, owner, hash, public) 
		VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id;`
	row, err := r.db.Exec(query,
		payload.Name,
		payload.MimeType,
		payload.Size,
		payload.SystemPath,
		payload.Owner,
		payload.Hash,
		payload.Public,
	)
	if err != nil {
		return 0, err
	}
	id, err = row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *DatabaseRepository) GetScreenListData(parent_id int, owner_id int) (*utils.FilesListResponse, error) {
	response := utils.FilesListResponse{}
	var rows *sql.Rows
	var err error
	if parent_id != 0 {
		query := `SELECT id, name, mime_type, size 
					FROM files 
					WHERE folder=? AND owner=?;`
		rows, err = r.db.Query(query, parent_id, owner_id)
		if err != nil {
			return nil, err
		}
	} else {
		query := `SELECT id, name, mime_type, size 
					FROM files 
					WHERE folder IS NULL AND owner=?;`
		rows, err = r.db.Query(query, owner_id)
		if err != nil {
			return nil, err
		}
	}
	for rows.Next() {
		file := utils.File{}
		err := rows.Scan(
			&file.Id,
			&file.Name,
			&file.MimeType,
			&file.Size,
		)
		if err != nil {
			return nil, err
		}
		response.Files = append(response.Files, file)
	}

	return &response, nil
}
