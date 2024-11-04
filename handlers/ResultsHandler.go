package handlers

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
)

type ResultsHandler struct {}

func (h ResultsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector(
		colly.AllowedDomains("https://jisho.org/"),
	)

	c.OnHTML(".concept-light.clearfix", func(h *colly.HTMLElement) {
		
	})

	fmt.Println(r.URL.Query().Get("search"))	
}