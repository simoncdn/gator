package main

import (
	"fmt"

	"github.com/simoncdn/gator/internal/config"
)

func main() {
	configFile, err := config.GetConfigFilePath();
	if err != nil {
		fmt.Printf("Error")
	}
	fmt.Println(configFile)
}
