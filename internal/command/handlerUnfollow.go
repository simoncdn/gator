package command

import (
	"context"
	"fmt"

	"github.com/simoncdn/gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	unfollowParams := database.UnfollowFeedParams {
		UserID: user.ID,
		Url: cmd.Args[0],
	}

	err := s.DB.UnfollowFeed(context.Background(), unfollowParams)
	if err != nil {
		return fmt.Errorf("couldn't unfollow this feed: %w", err)
	}
	
	fmt.Println("Unfollow this feed with success!")
	return nil
}
