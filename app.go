package main

import (
	"net/http"
)

func landingHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "landing", nil)
}

func choicesHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "choices", nil)
}