package rss

import "time"

type RssItem struct {
	Title       string
	Source      string
	SourceURL   string
	Link        string
	PublishDate time.Time
	Description string
}

func Parse() (rss RssItem, err error) {
	// implementation
	return
}
