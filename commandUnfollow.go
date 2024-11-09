package main

import (
	"context"
	"fmt"

	"github.com/mmammel12/aggreGATOR/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("'unfollow' command expects a feed url - Example: follow \"https://hnrss.org/newest\" ")
	}

	deleteFeedFolloParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    cmd.args[0],
	}
	err := s.db.DeleteFeedFollow(context.Background(), deleteFeedFolloParams)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully deleted follow for %v\n", cmd.args[0])

	return nil
}
