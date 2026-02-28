package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {

	db, err := sql.Open("sqlite", "pomodoro.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connetcted to database")

	todosTable := `
		CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT,
		session_id INTEGER,
		completed boolean,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP

		)
	`

	spotifyTable := `
		CREATE TABLE IF NOT EXISTS spotify (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		session_id INTEGER,
		access_token TEXT,
		refresh_token TEXT,
		expires_at INTEGER
		)
	`
	if _, err := db.Exec(todosTable); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(spotifyTable); err != nil {
		log.Fatal(err)
	}

	return db
}
