package main

import (
	"fmt"
	"time"
	"context"

	"github.com/borisfritz/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: follow <url>")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedsByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Unable to get feed: %w", err)
	}
	
	userID := user.ID
	feedID := feed.ID
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: userID,
		FeedID: feedID,
	})
	if err != nil {
		return fmt.Errorf("Unable to create Follow: %w", err)
	}
	fmt.Printf("User %v is now following %v", feedFollow.UserName, feedFollow.FeedName)
	return nil
} 
