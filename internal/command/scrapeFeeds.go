package command

import (
	"context"
	"fmt"

	"github.com/simoncdn/gator/internal/rss"
)

func scrapeFeeds(s *State) error {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't retreive this feed: %w", err)
	}
	
	_, err = s.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error on marking feed fetched: %w", err)
	}

	feedData, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("Couldn't collect feed %s: %v", feed.Name, err)
	}
	for i, item := range feedData.Channel.Item {
		fmt.Printf("- Feed %d\n", i)
		fmt.Printf("	* Title:               %s\n", item.Title)
		fmt.Println()
	}

	fmt.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
	return nil
}
