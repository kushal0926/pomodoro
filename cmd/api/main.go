package main

import (
	"encoding/gob"
	"log"

	"github.com/joho/godotenv"
	"github.com/kushal0926/pomodoro/internal/database"
	"github.com/kushal0926/pomodoro/internal/middleware"
	"github.com/kushal0926/pomodoro/internal/server"
)

func main() {
	godotenv.Load()
	middleware.InitStore()
	database.InitDB()
	gob.Register(int64(0))
	gob.Register(int(0))
	gob.Register(bool(false))

	srv := server.NewServer()
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
