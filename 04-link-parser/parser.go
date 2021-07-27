package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func collectNodeInnerText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	} else {
		var r string
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			r = r + collectNodeInnerText(c)
		}
		return r
	}
}

func ParseLinks(reader io.Reader) ([]Link, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	var results []Link

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			var href string
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href = attr.Val
				}
			}

			if href == "" || href == "#" {
				return
			}

			results = append(results, Link{
				Href: href,
				Text: strings.TrimSpace(collectNodeInnerText(n))})
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				walk(c)
			}
		}
	}
	walk(doc)

	return results, nil
}
