package handlers

import (
	"html/template"
	"net/http"
)

type BaseHandler struct {}

type BasePageData struct {
	LastSearch string
}

func (h BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"attr": func(attr string) template.HTMLAttr { return template.HTMLAttr(attr) },
	}
	tmpl := template.Must(template.Must(template.New("").Funcs(funcMap).ParseFiles("assets/html/base.html")).Parse(`{{define "content"}}{{end}}`))
	tmpl.ExecuteTemplate(w, "base", &BasePageData{LastSearch: ""})
}