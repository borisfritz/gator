package main

import (
	"fmt"
	"context"
)

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Failed to fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", rssFeed)
	return nil
}
