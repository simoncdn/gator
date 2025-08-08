package command

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/simoncdn/gator/internal/database"
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

	for _, item := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		newPost := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url:         item.Link,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		}

		_, err := s.DB.CreatePost(context.Background(), newPost)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			fmt.Printf("Couldn't create post: %v", err)
			continue
		}
	}

	fmt.Printf("Feed %s collected, %v posts found\n", feed.Name, len(feedData.Channel.Item))
	return nil
}
