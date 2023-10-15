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

	if config.Web.Port != 5000 {
		t.Errorf("Invalid parsing web port. Got: %d\n", config.Web.Port)
	}

	if config.Web.Host != "localhost" {
		t.Errorf("Invalid parsing web host. Got: %s\n", config.Web.Host)
	}

	if config.Web.PublicLink != "localhost:5000" {
		t.Errorf("Invalid parsing web public link. Got: %s\n", config.Web.PublicLink)
	}

	if config.Web.SessionName != "hfs_session" {
		t.Errorf("Invalid parsing web session name. Got: %s\n", config.Web.SessionName)
	}

	if config.Web.SecretKey != "test-secret-key" {
		t.Errorf("Invalid parsing web secret name. Got: %s\n", config.Web.SecretKey)
	}

	if config.CSRF.CsrfString != "test-csrf-string" {
		t.Errorf("Invalid parsing csrf string. Got: %s\n", config.CSRF.CsrfString)
	}
}
