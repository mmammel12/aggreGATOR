package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func handlerFeeds(s *state, _ command) error {
	feeds, err := s.db.ListFeedsWithUsers(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		data, err := json.MarshalIndent(feed, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(data))
	}

	return nil
}
