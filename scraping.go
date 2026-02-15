package main

import (
	"context"
	"fmt"
	"github.com/rQxwX3/gator/internal/rss"
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

	for _, rssItem := range rssFeed.Channel.Item {
		fmt.Println(rssItem.Title)
	}

	return nil
}
