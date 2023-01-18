package database

import (
	"gorm.io/gorm"
)

type createUserPayload struct {
	email           string
	hashed_password string
}

func GetUser(id int, db *gorm.DB) {
	db.Raw("SELECT id, email, active, created_at FROM users WHERE id=?;", id)
}

func CreateUser(payload createUserPayload, db *gorm.DB) {
	db.Raw("INSERT INTO users (email, hashed_password) VALUES (?, ?);", payload.email, payload.hashed_password)
}

// TODO
// 	- udpate
// 	- delete
