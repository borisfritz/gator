package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.name)
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Failed to set user: %w", err)
	}
	fmt.Println("User switched successfully!")
	return nil
}

