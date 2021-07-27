package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Web struct {
	story    Story
	template *template.Template
}

func NewWeb(story Story) Web {
	template, err := template.ParseFiles("page.gohtml")
	if err != nil {
		Exit(err)
	}
	return Web{story, template}
}

func (web Web) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pageId := strings.Trim(r.URL.Path, "/")
	if pageId == "" {
		pageId = "intro"
	}
	page, exists := web.story[pageId]
	if !exists {
		http.Error(w, "Unknown page", http.StatusNotFound)
		return
	}

	web.template.Execute(w, page)
}
