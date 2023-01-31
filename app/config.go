package app

import (
	"fmt"
	"homestorage/app/utils"

	"github.com/spf13/viper"
)

type ApplicationConfig struct {
	PORT int
}

type StorageConfig struct {
	SYSTEM_PATH string
}

type Config struct {
	Database    *utils.DBConfig
	Application *ApplicationConfig
	FileStorage *StorageConfig
}

func GetConfig(configPath string) Config {

	viper.SetConfigFile(configPath)

	var configuration Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration
}
