package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}
func Artist(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "artist.html")
}
func Error(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "error.html")
}
