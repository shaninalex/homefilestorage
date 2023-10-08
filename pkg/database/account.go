package database

type User struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	PasswordHash string `json:"-"`
}
