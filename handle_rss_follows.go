package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mate2xo/gator/internal/database"
)

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("please enter a <url> argument (usage: follow <url>)")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not find feed: %w", err)
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

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("choose an URL to unfollow (usage: unfollow <url>)")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not find feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return fmt.Errorf("could not delete feed_follow: %w", err)
	}

	fmt.Printf("%s feed is now unfollowed", feed.Name)
	return nil
}

func handlerFollowingFeeds(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not find feed_follows: %w", err)
	}

	if len(feedFollows) == 0 {
		println("No feeds followed.")
		return nil
	}

	println("Following:")
	for _, feedFollow := range feedFollows {
		fmt.Printf("- %s\n", feedFollow.FeedName)
	}
	return nil
}
