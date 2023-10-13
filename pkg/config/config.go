package config

import (
	"github.com/pelletier/go-toml"
)

// config reader from toml file
type Config struct {
	DB struct {
		Path string `toml:"path"`
	} `toml:"db"`
	Admin struct {
		Email    string `toml:"email"`
		Name     string `toml:"name"`
		Password string `toml:"password"`
	} `toml:"admin"`
	Storage struct {
		Path string `toml:"path"`
	} `toml:"storage"`
	GIN struct {
		Mode string `toml:"mode"`
		Port int64  `toml:"port"`
	} `toml:"gin"`
	Web struct {
		Port        int64  `toml:"port"`
		Host        string `toml:"host"`
		PublicLink  string `toml:"public_link"`
		SessionName string `toml:"session_name"`
	} `toml:"web"`
	CSRF struct {
		CsrfString string `toml:"csrf_string"`
	} `toml:"csrf"`
}

// TODO: tests for web port
func ParseConfig(config_path string) (*Config, error) {
	config, err := toml.LoadFile(config_path)
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := config.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
