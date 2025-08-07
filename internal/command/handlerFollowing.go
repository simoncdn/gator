package command

import (
	"context"
	"fmt"

	"github.com/simoncdn/gator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
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
