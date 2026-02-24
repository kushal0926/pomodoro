package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatalf("failed to load template: %v", err)
		}
		tmpl.Execute(w, nil)

	})

	http.ListenAndServe(":8080", nil)
}
