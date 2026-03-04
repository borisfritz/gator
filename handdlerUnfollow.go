package main

import (
	"context"
	"fmt"

	"github.com/borisfritz/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: unfollow <url>")
	}
	url := cmd.args[0]
	feed, err := s.db.GetFeedsByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Unable to retrieve feed: %w", err)
	}
	err = s.db.DeleteFeedFollowForUser(context.Background(), database.DeleteFeedFollowForUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Unable to delete feed: %w", err)
	}
	fmt.Printf("%v unfollowed '%v'", user.Name, feed.Url)
	return nil
}
