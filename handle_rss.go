package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Mate2xo/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("there should be 2 arguments (usage: addFeed <name> <url>)")
	}

	current_user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	name := cmd.args[0]
	url := cmd.args[1]
	fmt.Printf("name: %v, url: %v", name, url)
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    current_user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	fmt.Printf("Created feed: %+v", feed)
	return nil
}
