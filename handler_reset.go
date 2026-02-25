package main

import (
	"fmt"
	"context"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Usage: %s", cmd.name)
	}
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to delete users: %w", err)
	}
	return nil
}
