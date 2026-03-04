package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/borisfritz/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: agg <time_between_reqs>")
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Invalid duration: %w", err)
	}
	log.Printf("Collecting feeds every %s...\n", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("Unable to fetch next feed: %v\n", err)
		return
	}
	log.Println("Found feed to fetch!")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Failed to mark feed '%v' as fetched: %v\n", feed.Name, err)
		return
	}
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Println("Failed to fetch feed '%v': %v\n", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found Post: %v\n", item.Title)
	}
	log.Printf("Feed '%v' collected, %v posts found.", feed.Name, len(feedData.Channel.Item))
}
