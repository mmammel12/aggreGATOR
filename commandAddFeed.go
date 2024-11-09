package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/mmammel12/aggreGATOR/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("'addfeed' command expects a feed name and url - Example: addfeed \"Hacker News RSS\" \"https://hnrss.org/newest\" ")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("User %v not found in the database. Make sure you are registered", s.cfg.CurrentUserName)
	}

	feedParams := database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   cmd.args[0],
		Url:    cmd.args[1],
		UserID: user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
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
