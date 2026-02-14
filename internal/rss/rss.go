package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func unescapeRSSFeed(feed *RSSFeed) {
	html.UnescapeString(feed.Channel.Title)
	html.UnescapeString(feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		html.UnescapeString(item.Title)
		html.UnescapeString(item.Description)
	}
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")
	var client = &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var feed = &RSSFeed{}

	xml.Unmarshal(data, feed)
	unescapeRSSFeed(feed)
	return feed, nil
}
