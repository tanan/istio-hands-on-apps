package main

import (
	"html/template"
	"log"
	"net/http"
)

type NewsData struct {
	PageTitle    string
	NewsContents []NewsContent
}

type NewsContent struct {
	Title       string
	Description string
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/top.html"))
	data := NewsData{
		PageTitle: "top page",
		NewsContents: []NewsContent{
			{Title: "istio v1.0.0 release!!", Description: "this is a description statement"},
			{Title: "istio v1.0.1 release!!", Description: "this is a description statement"},
			{Title: "istio v1.0.2 release!!", Description: "this is a description statement"},
			{Title: "istio v1.0.3 release!!", Description: "this is a description statement"},
		},
	}
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/news", newsHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
