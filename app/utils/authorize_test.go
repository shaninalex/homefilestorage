package utils

import (
	"log"
	"testing"
)

// GenerateJWT
// RefreshJWT
// IdentifyJWT

var (
	user_email string = "test@test.com"
	user_id    int    = 1
)

func TestJWTGEneration(t *testing.T) {
	creds, err := GenerateJWT(user_email, user_id)
	if err != nil {
		t.Errorf("Cant generate JWT with given email %s and id %d", user_email, user_id)
	}

	email, id, err := IdentifyJWT(creds.Access)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Cant identify JWT with given email %s and id %d", user_email, user_id)
	}

	if email != &user_email && id != &user_id {
		t.Errorf("Emails and ids are not the same. email %s and id %d", user_email, user_id)
	}
}
