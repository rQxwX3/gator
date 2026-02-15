package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/rQxwX3/gator/internal/database"
	"github.com/rQxwX3/gator/internal/rss"
	"time"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	s.db.MarkFeedFetched(context.Background(), feed.ID)

	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	currentTime := time.Now()

	for _, rssItem := range rssFeed.Channel.Item {
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
			Title:       rssItem.Title,
			Url:         rssItem.Link,
			Description: rssItem.Description,
			PublishedAt: rssItem.PubDate,
			FeedID:      feed.ID,
		})

		if err != nil {
			if pqErr := err.(*pq.Error); pqErr.Code == "23505" {
				continue
			}

			return err
		}
	}

	return nil
}
