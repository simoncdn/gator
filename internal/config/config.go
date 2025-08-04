package config

import (
	"encoding/json"
	"os"
)

const (
	configFileName = ".gatorConfig.json"
)

type Config struct {
	Db_URL string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func getConfigFilePath() (string, error){
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := homePath + "/" + configFileName

	return fullPath, nil
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.Username = username
	err := write(*cfg)
	if err != nil {
		return err
	}
	return nil
}

func write (cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0666)
	if err != nil {
		return err
	}

	return nil
}
