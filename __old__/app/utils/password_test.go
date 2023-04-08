package utils

import (
	"log"
	"testing"
)

var test_password string = "test_password"

func TestComparingPassword(t *testing.T) {
	hashed_password, err := GenerateFromPassword(test_password)
	if err != nil {
		log.Fatal(err)
	}

	res, err := CheckPassword(test_password, hashed_password)
	if err != nil {
		log.Fatal(err)
	}

	if res != true {
		t.Errorf("Wrong password: %s and testing password: %s", hashed_password, test_password)
	}
}
