package main

import (
	"context"
	"encoding/xml"
	"fmt"
)

func handlerAgg(_ *state, _ command) error {
	url := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	data, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))

	return nil
}
