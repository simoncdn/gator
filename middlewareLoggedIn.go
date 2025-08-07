package main

import (
	"context"
	"fmt"

	"github.com/simoncdn/gator/internal/command"
	"github.com/simoncdn/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *command.State, cmd command.Command, user database.User) error) func(*command.State, command.Command) error {
	return func(s *command.State, cmd command.Command) error {
		user, err := s.DB.GetUserByName(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't find user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
