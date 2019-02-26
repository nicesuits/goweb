package main

import (
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)
		publicTemplateFiles := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html"}
		privateTemplateFiles := []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html"}
		var templates *template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(private_template_files...))
		} else {
			templates = template.Must(template.ParseFiles(public_template_files...))
		}
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
