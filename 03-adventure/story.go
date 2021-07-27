package main

import (
	"encoding/json"
	"io"
	"os"
)

type Page struct {
	Title       string
	Description []string `json:"story"`
	Options     []struct {
		Text string
		Arc  string
	}
}

type Story map[string]Page

func ImportStory(filename string) (Story, error) {
	var story Story

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &story)
	if err != nil {
		return nil, err
	}

	return story, nil
}
