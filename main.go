package main

import (
	"causeurgnocchi/dictionary/handlers"
	"net/http"
)


func main() {
	assets := http.FileServer(http.Dir("assets"))
	
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.Handle("/", &handlers.BaseHandler{})
	http.Handle("/results", http.StripPrefix("/results", &handlers.ResultsHandler{}))

	http.ListenAndServe(":8080", nil)
}
