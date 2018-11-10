package main

import (
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	PageTitle string
	Contents  []Content
}

type Content struct {
	Path  string
	Title string
}

func topHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/top.html"))
	data := Data{
		PageTitle: "top page",
		Contents: []Content{
			{Path: "news", Title: "news"},
			{Path: "blog", Title: "blog"},
		},
	}
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", topHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
