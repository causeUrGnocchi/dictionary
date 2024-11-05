package handlers

import (
	"html/template"
	"net/http"

	"github.com/gocolly/colly"
)

type ResultsHandler struct {}

type ResultsPageData struct {
    Words []word
    Search string
}

type word struct {
    Writing string `selector:".concept_light-representation > .text"`
    Reading string `selector:".concept_light-representation > .furigana"`
    Meanings []string `selector:"concept_light-meanings.medium-9.columns > .meaning-meaning"`
}

func (h ResultsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    words := make([]word, 0)
    
    c := colly.NewCollector()
    
	c.OnHTML("#primary", func(container *colly.HTMLElement) {
        container.ForEach(".concept_light.clearfix", func(i int, e *colly.HTMLElement) {
            v := &word{}
            e.Unmarshal(v)
            words = append(words, *v)
        })
	})

    c.Visit("https://jisho.org/search/" + r.URL.Query().Get("search"))

    tmpl := template.Must(template.ParseFiles("assets/html/home.html", "assets/html/results.html"))
    tmpl.ExecuteTemplate(w, "base", ResultsPageData{Words: words, Search: r.URL.Query().Get("search")})
}