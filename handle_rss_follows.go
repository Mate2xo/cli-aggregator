package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mate2xo/gator/internal/database"
)

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("please enter a <url> argument (usage: follow <url>)")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not find feed: %w", err)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not find user: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID, FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error when creating feed_follow: %w", err)
	}

	fmt.Printf("%s is now followed by %s\n", feed.Name, user.Name)
	return nil
}

func handlerFollowingFeeds(s *state, cmd command) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not find feed_follows: %w", err)
	}

	println("Following:")
	for _, feedFollow := range feedFollows {
		fmt.Printf("- %s\n", feedFollow.FeedName)
	}
	return nil
}
