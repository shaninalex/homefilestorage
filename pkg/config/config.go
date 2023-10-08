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
		Email string `toml:"email"`
		Name  string `toml:"name"`
	} `toml:"admin"`
	Storage struct {
		Path string `toml:"path"`
	} `toml:"storage"`
	GIN struct {
		Mode string `toml:"mode"`
		Port int64  `toml:"port"`
	} `toml:"gin"`
	Web struct {
		Port int64 `toml:"port"`
	} `toml:"web"`
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
