package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/mmammel12/aggreGATOR/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		intLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = intLimit
	}

	getPostsParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), getPostsParams)
	if err != nil {
		return err
	}

	for _, post := range posts {
		data, err := json.MarshalIndent(post, "", "  ")
		if err != nil {
			return err
		}

		fmt.Printf("%v\n", string(data))
	}

	return nil
}
