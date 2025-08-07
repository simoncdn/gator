package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/simoncdn/gator/internal/database"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get the feed: %w", err)
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Printf("Feed Follow created with success!\n")
	fmt.Printf(" * Name:          %s\n", feed.Name)
	fmt.Printf(" * User:          %s\n", user.Name)
	fmt.Println()
	fmt.Println("=====================================")
	return nil
}
