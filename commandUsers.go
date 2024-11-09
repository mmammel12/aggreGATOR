package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, _ command) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		return err
	}

	loggedInUserName := s.cfg.CurrentUserName

	for _, user := range users {
		fmt.Printf("* %v", user.Name)
		if user.Name == loggedInUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}

	return nil
}
