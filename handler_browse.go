package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/borisfritz/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		if givenLimit, err := strconv.Atoi(cmd.args[0]); err == nil {
			limit = givenLimit
		} else {
			return fmt.Errorf("Invalid limit: %w", err)
		}
	}
	posts, err := s.db.GetPostForUser(context.Background(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("Failed to retrieve posts for user: %w", err)
	}
	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	fmt.Println("=====================================")
	for _, post := range posts {
		fmt.Printf("'%s' on %s\n", post.FeedName, post.PublishedAt.Time.Format("Mon Jan 2"))
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}

