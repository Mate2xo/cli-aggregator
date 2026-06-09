package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	println("> resetting Users")
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not reset users: %w", err)
	}

	println("> resetting Feeds")
	err = s.db.ResetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not reset feeds: %w", err)
	}

	println("Database reset successful.")
	return nil
}
