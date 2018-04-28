package main

import (
	"html/template"
	"net/http"
)

// ReadHomePage display home page
func (h *dbHandler) ReadHomePage(w http.ResponseWriter, r *http.Request) {

	templates := template.New("layout")
	template.Must(templates.ParseFiles("templates/home.tmpl"))

	err := templates.ExecuteTemplate(w, "home.tmpl", nil)

	if err != nil {
		panic(err)
	}
}
