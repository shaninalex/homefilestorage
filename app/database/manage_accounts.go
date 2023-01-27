package database

import (
	"database/sql"
	"homestorage/app/utils"
	"log"
	"time"
)

type CreateUserPayload struct {
	Email          string
	HashedPassword string
}

type User struct {
	Id              int
	Email           string
	Hashed_password string
	Active          bool
	Created_at      *time.Time
}

func GetUser(id int, db *sql.DB) (*User, error) {
	var user User
	row := db.QueryRow("SELECT id, email, active, created_at FROM users WHERE id=?;", id)
	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Active,
		&user.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string, db *sql.DB) (User, error) {
	var user User
	db.QueryRow("SELECT * FROM users WHERE email=?;", email).Scan(&user)
	return user, nil
}

func CreateUser(payload *CreateUserPayload, db *sql.DB) error {
	encodedHash, err := utils.GenerateFromPassword(payload.HashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	query := "INSERT INTO users (email, hashed_password) VALUES (?, ?) RETURNING id, email, hashed_password, active, created_at;"
	_, err = db.Exec(query, payload.Email, encodedHash)
	if err != nil {
		return err
	}
	return nil
}

// TODO
// 	- udpate
// 	- delete
