package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("DB successfully reset")

	return nil
}
