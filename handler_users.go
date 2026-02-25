package main

import (
	"fmt"
	"errors"
	"context"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: users")
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to get user list: %w", err)
	}
	if len(users) == 0 {
		return errors.New("No users registered")
	}
	for _, user := range users {
		current := ""
		if user.Name == s.cfg.CurrentUserName {
			current = " (current)"
		}
		fmt.Printf("* %v%v\n", user.Name, current)
	}
	return nil
}


