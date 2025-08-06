package command

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {
	currentUser, err := s.DB.GetUserByName(context.Background(), s.Cfg.CurrentUserName) 
	if err != nil {
		return fmt.Errorf("couldn't get current user information: %w", err)
	}

	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("couldn't get follows for the current user: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No follows found.")
		return nil
	}

	fmt.Println("Feeds:")
	fmt.Println()
	for _, feed := range feeds {
		fmt.Printf("- %s\n", feed.FeedName)
	}
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}
