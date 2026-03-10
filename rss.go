package main

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

// RSSEntry represents a single RSS entry.
type RSSEntry struct {
	Title       string
	URL         string
	PublishedAt time.Time
}

func rssFeed(url string, count int) []RSSEntry {
	fp := gofeed.NewParser()

	var feed *gofeed.Feed
	var err error
	for attempt := 1; attempt <= 3; attempt++ {
		feed, err = fp.ParseURL(url)
		if err == nil {
			break
		}
		fmt.Printf("RSS fetch attempt %d/3 failed for %s: %v\n", attempt, url, err)
		if attempt < 3 {
			time.Sleep(time.Duration(attempt*2) * time.Second)
		}
	}
	if err != nil {
		fmt.Printf("RSS feed unavailable after 3 attempts, skipping: %s\n", url)
		return nil
	}

	var r []RSSEntry
	for _, v := range feed.Items {
		if v.PublishedParsed == nil {
			continue
		}
		r = append(r, RSSEntry{
			Title:       v.Title,
			URL:         v.Link,
			PublishedAt: *v.PublishedParsed,
		})
		if len(r) == count {
			break
		}
	}

	return r
}
