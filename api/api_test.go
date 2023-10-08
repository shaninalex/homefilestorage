package api

import (
	"database/sql"
	"testing"

	"github.com/shaninalex/homefilestorage/pkg/config"
	"github.com/shaninalex/homefilestorage/pkg/database"
)

func db_connect() *sql.DB {
	db, _ := sql.Open("sqlite3", ":memory:")
	return db
}

func TestInitAPI(t *testing.T) {
	db := db_connect()
	repo := database.InitSQLiteRepository(db)

	config, err := config.ParseConfig("../pkg/config/config.toml")
	if err != nil {
		t.Error("Unable to parse config")
	}

	api, err := CreateApi(repo, config)

	if err != nil {
		t.Errorf("Create API return error: %s\n", err.Error())
	}

	if api.database == nil {
		t.Error("Database should not be equal to nil")
	}
}
