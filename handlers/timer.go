package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kushal0926/pomodoro/middleware"
)

func GetTimer(w http.ResponseWriter, r *http.Request) {
	sessions := middleware.GetSession(w, r)

	// checking if running exists
	isRunningValue, ok := sessions.Values["is_running"]
	if !ok || isRunningValue == nil {
		fmt.Fprint(w, "25:00")
		return
	}

	startTime := sessions.Values["start_time"].(int64)
	duration := sessions.Values["duration"].(int)
	isRunning := isRunningValue.(bool)

	if !isRunning {
		fmt.Fprint(w, "25:00")
		return
	}

	remaining := duration - int(time.Now().Unix()-startTime)

	minutes := remaining / 60
	seconds := remaining % 60
	log.Println("session values:", sessions.Values)
	fmt.Fprintf(w, "%02d:%02d", minutes, seconds)
}

func StartTimer(w http.ResponseWriter, r *http.Request) {
	sessions := middleware.GetSession(w, r)
	start_time := time.Now().Unix()
	duration := 25 * 60
	is_running := true

	sessions.Values["start_time"] = start_time
	sessions.Values["duration"] = duration
	sessions.Values["is_running"] = is_running
	sessions.Save(r, w)
	log.Println("saved session:", sessions.Values)

	fmt.Fprint(w, "<h1>24:20</h1>")

}
