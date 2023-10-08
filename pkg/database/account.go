package database

type Account struct {
	ID           int64  `json:"-" db:"id"`
	Email        string `json:"email" db:"email"`
	Name         string `json:"name" db:"name"`
	PasswordHash string `json:"-" db:"password_hash"`
}
