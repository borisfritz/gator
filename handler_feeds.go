package main

import (
	"fmt"
	"context"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: feeds")
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to get feed: %w", err)
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Unable to get User by ID: %w", err)
		}
		fmt.Printf("* Name: %v\n", feed.Name)
		fmt.Printf("* URL: %v\n", feed.Url)
		fmt.Printf("* User: %v\n", user.Name)
	}
	return nil
}
