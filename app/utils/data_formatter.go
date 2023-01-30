package utils

import (
	"time"
)

type PublicUserObject struct {
	Id         int        `json:"id"`
	Email      string     `json:"email"`
	Active     bool       `json:"active"`
	Created_at *time.Time `json:"created_at"`
}

func PublicUser(user *User) *PublicUserObject {
	// We need to remove hashed password from public available
	// user object
	public_user := &PublicUserObject{
		Id:         user.Id,
		Email:      user.Email,
		Active:     user.Active,
		Created_at: user.Created_at,
	}
	return public_user
}
