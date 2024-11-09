package main

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mmammel12/aggreGATOR/internal/database"
)

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	markFeedFetchedParams := database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}
	err = s.db.MarkFeedFetched(ctx, markFeedFetchedParams)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		err := description.Scan(item.Description)
		if err != nil {
			return err
		}

		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return err
		}

		createPostParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         feed.Url,
			Description: description,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		}
		_, err = s.db.CreatePost(ctx, createPostParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			return err
		}
	}

	return nil
}
