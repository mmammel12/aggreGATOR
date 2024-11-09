package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mmammel12/aggreGATOR/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("'follow' command expects a feed url - Example: follow \"https://hnrss.org/newest\" ")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
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
	follow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %v\nUser: %v\n", follow.FeedName, follow.UserName)

	return nil
}
