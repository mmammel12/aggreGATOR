package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, _ command) error {
	duration := 1 * time.Minute
	ticker := time.NewTicker(duration)

	fmt.Printf("Collecting feeds every %v\n", duration.String())

	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}
