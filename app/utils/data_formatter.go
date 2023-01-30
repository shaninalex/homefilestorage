package utils

import "homestorage/app/database"

func PublicUser(user *database.User) *database.User {
	// We need to remove hashed password from public available
	// user object
	user.Hashed_password = ""
	return user
}
