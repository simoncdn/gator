package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/simoncdn/gator/internal/database"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]
	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      username,
	}

	user, err := s.DB.CreateUser(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User created with success!")
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)

	return nil
}
