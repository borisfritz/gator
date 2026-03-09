package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"strings"
	"database/sql"

	"github.com/borisfritz/gator/internal/database"
	"github.com/google/uuid"
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
		log.Printf("Failed to mark feed '%v' as fetched: %v\n", feed.Name, err)
		return
	}
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Failed to fetch feed '%v': %v\n", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		// Check if description and publish-date exist and set value
		description := sql.NullString{
			String: item.Description,
			Valid: item.Description != "",
		}
		publishedAt := sql.NullTime{
			Valid: false,
		}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt.Time = t
			publishedAt.Valid = true
		}

		// Create post
		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: item.Title,
			Url: item.Link,
			Description: description,
			PublishedAt: publishedAt,
			FeedID: feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
				continue
			}
			log.Printf("Failed to create post '%v': %v", item.Title, err)
		}
	}
}
