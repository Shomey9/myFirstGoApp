package scraper

import (
	"fmt"
	"sync"

	"github.com/Shomey9/myFirstGoApp/internal/rss"
)

// import {
//     "fmt"
//     "sync"

//     "github.com/Ba5ik7/go-rss-feed/internal/rss"
// }

func ScrapeFeeds(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(feedURL string) {
			defer wg.Done()
			feed, err := rss.FetchFeed(feedURL)
			if err != nil {
				fmt.Printf("Error scraping %s: %v\n", feedURL, err)
				return
			}
			fmt.Printf("Feed Title: %s\n", feed.Channel.Title)
			for _, item := range feed.Channel.Items {
				fmt.Printf("- %s (%s)\n", item.Title, item.Link)
			}
		}(url)
	}

	wg.Wait()
}
