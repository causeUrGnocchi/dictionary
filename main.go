package main

import (
	"causeurgnocchi/dictionary/handlers"
	"html/template"
	"net/http"
)

type PageData struct {
	Search string
}

func main() {
	assets := http.FileServer(http.Dir("assets"))
	
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.Must(template.ParseFiles("assets/html/base.html")).Parse(`{{define "content"}}{{end}}`))
		tmpl.ExecuteTemplate(w, "base", &PageData{Search: ""})
	})

	http.Handle("/results", http.StripPrefix("/results", &handlers.ResultsHandler{}))

	http.ListenAndServe(":8080", nil)
}
