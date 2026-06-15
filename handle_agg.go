package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Mate2xo/gator/internal/database"
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
		log.Printf("Could not get next feed to fetch: %v", err)
		return
	}
	log.Println("Found feed to fetch")

	rssFeed, err := scrapeFeed(s, nextFeed)
	if err != nil {
		log.Printf("could not scrape feed: %v", err)
		return
	}
	err = savePosts(s, rssFeed, nextFeed)
	if err != nil {
		log.Printf("error: could not save scraped feed: %v", err)
	}
}

func scrapeFeed(s *state, feed database.Feed) (*RSSFeed, error) {
	remoteFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return nil, fmt.Errorf("could not fetch feed: %w", err)
	}
	err = s.db.MakeFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return nil, fmt.Errorf("could not mark feed: %w", err)
	}

	return remoteFeed, nil
}

func savePosts(s *state, rss *RSSFeed, feed database.Feed) error {
	fmt.Printf("--Link is %+v\n", rss.Channel.Link)
	for _, item := range rss.Channel.Item {
		parsedPubDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
		if err != nil {
			return fmt.Errorf("error parsing Item's PubDate: %w", err)
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description},
			PublishedAt: sql.NullTime{Time: parsedPubDate},
			FeedID:      feed.ID,
		})
		if err != nil {
			return fmt.Errorf("error creating post: %w (%T, %+v)", err, err, err)
		}
	}

	return nil
}
