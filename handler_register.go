package main

import (
	"fmt"
	"context"
	"time"

	"github.com/borisfritz/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.name)
	}
	name := cmd.args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: name, 
	})
	if err != nil {
		return fmt.Errorf("Failed to create User: %w", err)
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("Unable to set user in config file: %w", err)
	}
	fmt.Printf("%+v\n", user)
	return nil
}
