package config

import (
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error){
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homePath, configFileName)

	return fullPath, nil
}

