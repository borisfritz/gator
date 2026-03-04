package main

import (
	"context"
	"fmt"

	"github.com/borisfritz/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: following")
	}
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Unable to get followed feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Printf("%v is not following any feeds.", user.Name)
		return nil
	}

	fmt.Printf("User %v follows:\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf("* %v\n", feed.FeedName)
	}
	return nil
}
