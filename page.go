package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Page struct {
	ID      int    `db:"id" json:"id"`
	Slug    string `db:"slug" json:"slug"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
}

// ReadPage display post
func (h *dbHandler) ReadPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	var page Page
	err := h.db.QueryRowx("SELECT * FROM pages WHERE slug=?", slug).StructScan(&page)

	// var templates = template.Must(template.ParseFiles("templates/layout.tmpl", "templates/404.tmpl"))

	// if err != nil {
	// 	templates.ExecuteTemplate(w, "404", nil)
	// }

	templates := template.New("layout")
	template.Must(templates.ParseFiles("templates/layout.tmpl", "templates/404.tmpl"))

	if err != nil {
		templates.ExecuteTemplate(w, "404", nil)
	} else {

		template.Must(templates.New("content").Parse(page.Content))

		err = templates.ExecuteTemplate(w, "layout.tmpl", &page)

		if err != nil {
			panic(err)
		}
	}
}
