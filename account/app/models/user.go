package models

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type UpdateUser struct {
	Email    *string
	Username *string
}

type User struct {
	ID        int64     `json:"id,omitempty"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Active    bool      `json:"active"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func GetUser(db *sql.DB, userID int64) (*User, error) {
	var user User
	row := db.QueryRow(`SELECT * FROM users WHERE id = $1`, userID)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Active,
		&user.Password,
		&user.UpdatedAt,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Create(db *sql.DB) (int64, error) {
	err := db.QueryRow(
		`INSERT INTO users (email, username, active, password) VALUES ($1, $2, $3, $4) RETURNING id;`,
		u.Email, u.Username, u.Active, u.Password,
	).Scan(&u.ID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return u.ID, nil
}

func (u *User) Update(db *sql.DB) error {
	res, err := db.Exec(
		`UPDATE users SET email = $1, Username = $2 WHERE id = $3`,
		u.Email, u.Username, u.ID,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	id, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	if id == 0 {
		err = errors.New("no rows effected")
		return err
	}
	return nil
}

func (u *User) Delete(db *sql.DB, ID int64) error {
	res, err := db.Exec(`DELETE FROM users WHERE id = $1`, ID)
	if err != nil {
		log.Println(err)
		return err
	}
	id, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	if id == 0 {
		err = errors.New("no rows effected")
		return err
	}
	return nil
}
