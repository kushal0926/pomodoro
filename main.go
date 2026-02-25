package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/kushal0926/pomodoro/db"
	"github.com/kushal0926/pomodoro/middleware"
)

func main() {
	db.InitDB()
	fmt.Println("making an pomodoro app")
	// main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatalf("failed to load template: %v", err)
		}
		tmpl.Execute(w, nil)
		session :=middleware.GetSession(w, r)
		log.Println(session)

	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>helllo from go </h1>")
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		formatted := currentTime.Format("2006-01-02 15:04:05")
		fmt.Fprint(w, formatted)
	})

	http.ListenAndServe(":8080", nil)
}
