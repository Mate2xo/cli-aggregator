package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 || len(cmd.args) > 2 {
		return fmt.Errorf("please enter a duration (usage: %s <time between requests>)", cmd.name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("Could not get next feed to fetch", err)
		return
	}
	log.Println("Found feed to fetch")

	updatedFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		log.Printf("could not fetch feed: %w", err)
		return
	}
	err = s.db.MakeFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		log.Printf("could not mark feed: %w", err)
		return
	}

	fmt.Printf("Collected feed titles from %s:", nextFeed.Name)
	for _, item := range updatedFeed.Channel.Item {
		fmt.Printf("- %s\n", item.Title)
	}
}
