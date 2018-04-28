package main

import "net/http"

type Post struct {
	ID      int    `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
	Author  string `db:"author" json:"author"`
}

// ReadPost display post
func ReadPost(w http.ResponseWriter, r *http.Request) {

}
