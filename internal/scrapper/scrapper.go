package scrapper

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Kaivv1/blog-aggregator/internal/database"
)

func StartScrapping(
	db *database.Queries,
	concurency int,
	timeBetweemRequest time.Duration,
) {
	log.Printf("Scrapping on %v goroutines every %v duration", concurency, timeBetweemRequest)

	ticker := time.NewTicker(timeBetweemRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurency))
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error marking feed as fetched", err)
	}
	rssFeed, err := UrlToFeed(feed.Url)
	if err != nil {
		log.Println("error fetching feed", err)
		return
	}
}
