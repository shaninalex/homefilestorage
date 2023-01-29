package database

import (
	"database/sql"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// TODO
// 	- udpate
// 	- delete
//  - return predefined errors

func CreateDatabaseConnection(conf *DBConfig) (*DatabaseRepository, error) {

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

func (r *DatabaseRepository) GetUser(id int) (*User, error) {
	var user User
	row := r.db.QueryRow("SELECT id, email, active, hashed_password, created_at FROM users WHERE id=?;", id)
	err := row.Scan(&user.Id, &user.Email, &user.Active, &user.Hashed_password, &user.Created_at)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *DatabaseRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	row := r.db.QueryRow("SELECT id, email, active, hashed_password, created_at FROM users WHERE email=?;", email)
	err := row.Scan(&user.Id, &user.Email, &user.Active, &user.Hashed_password, &user.Created_at)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *DatabaseRepository) CreateUser(payload *CreateUserPayload) error {
	// TODO: Remove generating password hash from database package!
	query := "INSERT INTO users (email, hashed_password) VALUES (?, ?) RETURNING id;"
	_, err := r.db.Exec(query, payload.Email, payload.HashedPassword)
	if err != nil {
		return err
	}
	return nil
}
