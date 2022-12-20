package main

import (
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{
		"Dayane",
		"day@email.com",
	}
	templates.ExecuteTemplate(w, "home.html", u)
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
