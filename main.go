package main

import (
	"fmt"

	"github.com/simoncdn/gator/internal/config"
)

func main() {
	configFile, err := config.Read();
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Println(configFile)
}
