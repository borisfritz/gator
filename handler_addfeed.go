package main

import (
	"fmt"
	"time"
	"context"

	"github.com/borisfritz/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("Usage: addfeed <name> <url>")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("Failed to create feed: %w", err)
	}
	fmt.Println("Feed Created!")
	printFeed(feed)
	fmt.Println()
	fmt.Println("================================")
	fmt.Println()

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Unable to create Follow: %w", err)
	}
	fmt.Printf("User %v is now following %v", feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:       %s\n", feed.ID)
	fmt.Printf("* Created:  %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:  %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:     %s\n", feed.Name)
	fmt.Printf("* URL:      %s\n", feed.Url)
	fmt.Printf("* UserID:   %s\n", feed.UserID)
}
