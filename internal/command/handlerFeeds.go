package command

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}


	fmt.Println("List of feeds:")
	fmt.Println()
	for i, feed := range feeds {
		user, err := s.DB.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("fetch user error: %w", err)
		}
		fmt.Printf("- Feed %d\n", i + 1)
		fmt.Printf(" * Name:          %s\n", feed.Name)
		fmt.Printf(" * URL:           %s\n", feed.Url)
		fmt.Printf(" * User:          %s\n", user.Name)
		fmt.Println()
	}
	fmt.Println("=====================================")
	return nil
}
