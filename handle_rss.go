package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mate2xo/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("there should be 2 arguments (usage: addFeed <name> <url>)")
	}

	name := cmd.args[0]
	url := cmd.args[1]
	fmt.Printf("name: %v, url: %v\n", name, url)
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}
	fmt.Printf("Created feed: %+v\n", feed)

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID, FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not follow feed: %w", err)
	}
	fmt.Printf("%s is now followed by %s\n", feed.Name, user.Name)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not retrieve feeds: %w", err)
	}

	for _, feed := range feeds {
		feedUser, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("could not get user: %w", err)
		}
		fmt.Printf("- %s (%s) -- %s\n", feed.Name, feed.Url, feedUser.Name)
	}
	return nil
}
