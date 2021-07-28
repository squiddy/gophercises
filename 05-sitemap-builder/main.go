package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	parser "reinergerecke.de/gophercises/04-link-parser"
)

func isRelevantLink(origin *url.URL, target *url.URL) bool {
	if origin.Host != target.Host {
		return false
	}

	if origin.String() == target.String() {
		return false
	}

	return true
}

func getLinksFromPage(base string) ([]string, error) {
	var results []string

	baseUrl, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(base)
	if err != nil {
		return nil, err
	}

	links, err := parser.ParseLinks(res.Body)
	if err != nil {
		return nil, err
	}

	for _, link := range links {
		url, err := url.Parse(link.Href)
		if err != nil {
			return nil, err
		}

		// Normalize URL
		if url.Path == "" {
			url.Path = "/"
		}
		url.RawQuery = ""
		url.Fragment = ""

		resolved := baseUrl.ResolveReference(url)
		if isRelevantLink(baseUrl, resolved) {
			results = append(results, resolved.String())
		}
	}

	return results, nil
}

func getLinks(target *url.URL, depth int) ([]string, error) {
	links := make(map[string]struct{})
	var queue map[string]struct{}
	nextQueue := map[string]struct{}{
		target.String(): {},
	}

	for i := 0; i < depth; i++ {
		queue, nextQueue = nextQueue, make(map[string]struct{})
		for link := range queue {
			links[link] = struct{}{}
			links, err := getLinksFromPage(link)
			if err != nil {
				return nil, err
			}
			for _, l := range links {
				nextQueue[l] = struct{}{}
			}

		}
	}

	var result []string
	for url := range links {
		result = append(result, url)
	}
	return result, nil
}

type Url struct {
	Loc string `xml:"loc"`
}

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	XmlNS   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

func main() {
	depth := flag.Int("depth", 3, "")
	target := flag.String("target", "", "")
	flag.Parse()

	targetUrl, err := url.Parse(*target)
	if err != nil {
		Exit(err)
	}

	links, err := getLinks(targetUrl, *depth)
	if err != nil {
		Exit(err)
	}

	sitemap := Urlset{XmlNS: "http://www.sitemaps.org/schemas/sitemap/0.9", Urls: []Url{}}
	for _, link := range links {
		sitemap.Urls = append(sitemap.Urls, Url{Loc: link})
	}

	data, err := xml.Marshal(sitemap)
	if err != nil {
		Exit(err)
	}
	fmt.Print(xml.Header)
	fmt.Print(string(data))
}

func Exit(err error) {
	log.Print(err)
	os.Exit(1)
}
