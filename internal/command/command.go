package command

import (
	"errors"
	"fmt"

	"github.com/simoncdn/gator/internal/config"
	"github.com/simoncdn/gator/internal/database"
)

type State struct {
	Cfg *config.Config
	DB  *database.Queries
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CommandsList map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	command, ok := c.CommandsList[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}

	return command(s, cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.CommandsList[name] = f
	fmt.Printf("New commande register: %v\n", name)
}
