package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mmammel12/aggreGATOR/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("'addfeed' command expects a feed name and url - Example: addfeed \"Hacker News RSS\" \"https://hnrss.org/newest\" ")
	}

	feedParams := database.CreateFeedParams{
		ID:            uuid.New(),
		Name:          cmd.args[0],
		Url:           cmd.args[1],
		UserID:        user.ID,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		LastFetchedAt: sql.NullTime{},
	}
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(feed, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}
