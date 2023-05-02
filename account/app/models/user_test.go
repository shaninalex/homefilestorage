package models

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestUserCRUD(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer cleanDB(db)

	// Create
	u1 := &User{
		Email:    "user1@example.com",
		Username: "user1",
		Active:   true,
		Password: "password1",
	}
	id, err := u1.Create(db)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Error("Create failed to return ID")
	}
	if u1.ID == 0 {
		t.Error("Create failed to set ID")
	}

	// Read
	u2, err := GetUser(db, u1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if u2.Email != u1.Email {
		t.Errorf("Expected email '%s', got '%s'", u1.Email, u2.Email)
	}
	if u2.Username != u1.Username {
		t.Errorf("Expected username '%s', got '%s'", u1.Username, u2.Username)
	}
	if u2.Active != u1.Active {
		t.Errorf("Expected active '%v', got '%v'", u1.Active, u2.Active)
	}
	if u2.Password != u1.Password {
		t.Errorf("Expected password '%s', got '%s'", u1.Password, u2.Password)
	}

	// Update
	update := &UpdateUser{
		Email:    strPtr("new-email@example.com"),
		Username: strPtr("new-username"),
	}
	u1.Email = *update.Email
	u1.Username = *update.Username
	err = u1.Update(db)
	if err != nil {
		t.Fatal(err)
	}
	u3, err := GetUser(db, u1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if u3.Email != u1.Email {
		t.Errorf("Expected email '%s', got '%s'", u1.Email, u3.Email)
	}
	if u3.Username != u1.Username {
		t.Errorf("Expected username '%s', got '%s'", u1.Username, u3.Username)
	}

	// Delete
	err = u1.Delete(db, u1.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = GetUser(db, u1.ID)
	if err == nil {
		t.Error("User was not deleted")
	}
}

func strPtr(s string) *string {
	return &s
}

func cleanDB(db *sql.DB) {
	db.Exec(`DELETE FROM users`)
	db.Close()
}
