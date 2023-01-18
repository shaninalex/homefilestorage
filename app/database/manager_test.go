package database

import "testing"

var conf DBConfig = DBConfig{
	DEBUG:      true,
	DEBUG_NAME: "database.db",
	HOST:       "localhost",
	PORT:       5432,
	NAME:       "postgres",
	USER:       "postgres",
	PASS:       "postgres",
}

func TestAdd(t *testing.T) {

	got := buildConnectionUrl(&conf)
	want := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
