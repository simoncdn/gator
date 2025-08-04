package main

import (
	"fmt"
	"log"
	"os"

	"github.com/simoncdn/gator/internal/command"
	"github.com/simoncdn/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
	}

	programState := &command.State{
		Cfg: &cfg,
	}

	commands := command.Commands{
		CommandsList: map[string]func(*command.State, command.Command) error {},
	}
	commands.Register("login", command.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Printf("argumenst error\n")
		os.Exit(1)
	}

	cmd := command.Command {
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = commands.Run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
