package config

import (
	"os"
)

const (
	configFileName = ".gatorConfig.json"
)

type Config struct {
	Db_URL string `json:"db_url"`
}

func GetConfigFilePath() (string, error){
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := homePath + "/" + configFileName

	return fullPath, nil
}

func (cfg *Config) Read() (Config, error) {
	return Config{}, nil
}

