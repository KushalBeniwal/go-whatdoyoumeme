package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("./web/templates/landing.html", "./web/templates/choices.html"))

func main() {
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./web/static/assets/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/static/stylesheets/"))))

	http.HandleFunc("/", landingHandler)
	http.HandleFunc("/choices", choicesHandler)
	http.ListenAndServe(":8080", nil)
}