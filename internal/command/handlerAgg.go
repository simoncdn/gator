package command

import (
	"context"
	"fmt"

	"github.com/simoncdn/gator/internal/rss"
)

const (
	defaultRSSURL = "https://www.wagslane.dev/index.xml"
)

func HandlerAgg(s *State, cmd Command) error {
	feed, err := rss.FetchFeed(context.Background(), defaultRSSURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	fmt.Println(feed)
	return nil
}
