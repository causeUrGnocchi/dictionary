package main

import (
	"net/http"

	"github.com/gocolly/colly"
)

type vocabulary struct {
	Reading string `selector:".concept_light-representation > .text"`
	Meanings []string `selector:"span.meaning-meaning"`
}

func main() {
	c := colly.NewCollector(colly.AllowedDomains("jisho.org"))

	c.OnHTML("html", func(e *colly.HTMLElement) {
		e.ForEach(".concept_light.clearfix", func(i int, vocabularyElement *colly.HTMLElement) {

		})
		
		v := &vocabulary{}
		e.Unmarshal(v)
	})

	c.Visit("https://jisho.org/search/%E5%B1%95%E9%96%8B")

	assets := http.FileServer(http.Dir("assets"))
	
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/html/home.html")
	})

	http.ListenAndServe(":8080", nil)
}
