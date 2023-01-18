package database

import (
	"gorm.io/gorm"
)

type CreateUserPayload struct {
	Email          string
	HashedPassword string
}

func GetUser(id int, db *gorm.DB) {
	db.Raw("SELECT id, email, active, created_at FROM users WHERE id=?;", id)
}

func CreateUser(payload *CreateUserPayload, db *gorm.DB) {
	db.Raw("INSERT INTO users (email, hashed_password) VALUES (?, ?);", payload.Email, payload.HashedPassword)
}

// TODO
// 	- udpate
// 	- delete
