package config

import (
	"testing"
)

func TestConfigParser(t *testing.T) {
	config, err := ParseConfig("./config.toml")
	if err != nil {
		t.Errorf("ParseConfig return error: %s\n", err.Error())
	}

	if config.DB.Path != "database.db" {
		t.Errorf("Invalid parsing database path. Got: %s\n", config.DB.Path)
	}

	if config.Admin.Email != "email@email.com" {
		t.Errorf("Invalid parsing admin email. Got: %s\n", config.Admin.Email)
	}

	if config.Admin.Name != "Test Name" {
		t.Errorf("Invalid parsing admin name. Got: %s\n", config.Admin.Name)
	}

	if config.Storage.Path != "./storage/" {
		t.Errorf("Invalid parsing storage path. Got: %s\n", config.Storage.Path)
	}

	if config.GIN.Mode != "release" {
		t.Errorf("Invalid parsing gin mode config. Got: %s\n", config.GIN.Mode)
	}

	// TODO: gin mode can be only "release", "debug" or "test"

	if config.GIN.Port != 8080 {
		t.Errorf("Invalid parsing gin port config. Got: %d\n", config.GIN.Port)
	}
}
