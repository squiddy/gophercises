package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	story, err := ImportStory("gopher.json")
	if err != nil {
		Exit(err)
	}

	web := NewWeb(story)
	http.ListenAndServe(":8000", web)
}

func Exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
