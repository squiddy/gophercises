package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Config []struct {
	Path string
	Url  string
}

func main() {
	configName := flag.String("config", "config.yml", "")

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	}))

	config, err := readConfig(*configName)
	if err != nil {
		exit(err)
	}

	for _, rule := range config {
		mux.Handle(rule.Path, http.RedirectHandler(rule.Url, http.StatusFound))
	}

	http.ListenAndServe(":8000", mux)
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func readConfig(filename string) (Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
