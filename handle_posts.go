package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Mate2xo/gator/internal/database"
	"github.com/google/uuid"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int
	var err error
	if len(cmd.args) == 1 {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			limit = 2
		}
	} else {
		limit = 2
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not fetch user feeds: %w", err)
	}
	feedIds := make([]uuid.UUID, len(feeds))
	for i, feed := range feeds {
		feedIds[i] = feed.FeedID
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("could not fetch user posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Printf("%v, from %s", post.PublishedAt.Time.Format("Mon 2 Jan"), post.Name)
		fmt.Printf("%v\n", post.Description.String)
		println("==================================================")
		println()
	}
	return nil
}
