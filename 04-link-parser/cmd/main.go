package main

import (
	"flag"
	"fmt"
	"os"

	parser "github.com/squiddy/gophercises/04-link-parser"
)

func main() {
	filename := flag.String("filename", "", "file to parse")
	flag.Parse()

	file, _ := os.Open(*filename)
	links, _ := parser.ParseLinks(file)
	for _, link := range links {
		fmt.Printf("[%s] %s\n", link.Href, link.Text)
	}
}
