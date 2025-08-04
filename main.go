package main

import (
	"fmt"

	"github.com/simoncdn/gator/internal/config"
)

func main() {
	configFile, err := config.Read();
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
	}

	fmt.Println("Initial config:", configFile)

	err = configFile.SetUser("Simon")
	if err != nil {
		fmt.Printf("Error setting user: %v\n", err)
	}

	configFile, err = config.Read()
	if err != nil {
		fmt.Printf("Error reading update config: %v\n", err)
	}

	fmt.Println("Updated config:", configFile)
}
