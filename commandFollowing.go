package main

import (
	"context"
	"fmt"

	"github.com/mmammel12/aggreGATOR/internal/database"
)

func handlerFollowing(s *state, _ command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Printf("Feed: %v\n", follow.FeedName)
	}

	return nil
}
