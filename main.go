package main

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/kushal0926/pomodoro/db"
	"github.com/kushal0926/pomodoro/handlers"
	"github.com/kushal0926/pomodoro/middleware"
)

func main() {
	godotenv.Load()
	middleware.InitStore()
	db.InitDB()
	fmt.Println("making an pomodoro app")
	gob.Register(int64(0))
	gob.Register(int(0))
	gob.Register(bool(false))

	// main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatalf("failed to load template: %v", err)
		}
		session := middleware.GetSession(w, r)
		tmpl.Execute(w, nil)
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

	http.HandleFunc("/timer", handlers.GetTimer)
	http.HandleFunc("/timer/start", handlers.StartTimer)
	http.HandleFunc("/timer/stop", handlers.StopTimer)

	http.ListenAndServe(":8080", nil)
}
