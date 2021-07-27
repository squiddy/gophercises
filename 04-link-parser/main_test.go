package main

import (
	"os"
	"reflect"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		file     string
		expected []Link
	}{
		{"ex1.html", []Link{
			{Href: "/other-page", Text: "A link to another page"},
		}},
		{"ex2.html", []Link{
			{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
			{Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"},
		}},
		{"ex3.html", []Link{
			{Href: "/lost", Text: "Lost? Need help?"},
			{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
		}},
		{"ex4.html", []Link{
			{Href: "/dog-cat", Text: "dog cat"},
		}},
	}

	for _, c := range cases {
		file, err := os.Open(c.file)
		if err != nil {
			t.Error(err)
		}

		links, err := ParseLinks(file)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(links, c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, links)
		}
	}
}
