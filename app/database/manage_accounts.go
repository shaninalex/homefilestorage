package database

import (
	"homestorage/app/utils"
	"log"

	"gorm.io/gorm"
)

type CreateUserPayload struct {
	Email          string
	HashedPassword string
}

func GetUser(id int, db *gorm.DB) {
	db.Exec("SELECT id, email, active, created_at FROM users WHERE id=?;", id)
}

func CreateUser(payload *CreateUserPayload, db *gorm.DB) {
	encodedHash, err := utils.GenerateFromPassword(payload.HashedPassword)

	if err != nil {
		log.Fatal(err)
	}
	db.Exec("INSERT INTO users (email, hashed_password) VALUES (?, ?);", payload.Email, encodedHash)
}

// TODO
// 	- udpate
// 	- delete
