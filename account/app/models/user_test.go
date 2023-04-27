package models

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserCRUD(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	// Run migrations
	db.AutoMigrate(&User{})

	// Create a new user
	user := &User{
		Email:    "test@example.com",
		Username: "testuser",
		Active:   true,
		Password: "password",
	}
	userID, err := user.Create(db)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}
	if userID == 0 {
		t.Fatalf("failed to get user ID")
	}

	// Get the user by ID
	gotUser, err := GetUser(db, userID)
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}
	if gotUser.ID != userID {
		t.Fatalf("got user with ID %d, want %d", gotUser.ID, userID)
	}

	// Update the user
	user.Username = "newusername"
	user.Active = false
	err = user.Update(db)
	if err != nil {
		t.Fatalf("failed to update user: %v", err)
	}

	// Get the user again to ensure the update was successful
	gotUser, err = GetUser(db, userID)
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}
	if gotUser.Username != user.Username {
		t.Fatalf("got username %s, want %s", gotUser.Username, user.Username)
	}
	if gotUser.Active != user.Active {
		t.Fatalf("got active %t, want %t", gotUser.Active, user.Active)
	}

	// Delete the user
	err = user.Delete(db, userID)
	if err != nil {
		t.Fatalf("failed to delete user: %v", err)
	}

	// Try to get the user again to ensure it was deleted
	_, err = GetUser(db, userID)
	if err == nil {
		t.Fatalf("expected an error when getting deleted user, got nil")
	}
	if err != gorm.ErrRecordNotFound {
		t.Fatalf("got error %v, want %v", err, gorm.ErrRecordNotFound)
	}
}
