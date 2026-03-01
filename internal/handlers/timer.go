package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kushal0926/pomodoro/internal/middleware"
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
	if remaining <= 0 {
		mode, ok := sessions.Values["mode"].(string)
		if !ok {
			mode = "work"
		}

		var newDuration int
		if mode == "work" {
			mode = "break"
			breakDur, ok := sessions.Values["break_duration"].(int)
			if !ok {
				breakDur = 5
			}
			newDuration = breakDur * 60

		} else {
			mode = "work"
			workDur, ok := sessions.Values["work_duration"].(int)
			if !ok {
				workDur = 35
			}
			newDuration = workDur * 60
		}

		sessions.Values["mode"] = mode
		sessions.Values["start_time"] = time.Now().Unix()
		sessions.Values["duration"] = newDuration
		sessions.Save(r, w)

		var message string
		if mode == "break" {
			message = "Take a break!"
		} else {
			message = "Back to work!"
		}
		fmt.Fprintf(w, `<div>00:00</div><div id="notify" hx-swap-oob="true"><script>playSound(); showPopup('%s');</script></div>`, message)
		return

	}

	minutes := remaining / 60
	seconds := remaining % 60
	log.Println("session values:", sessions.Values)
	fmt.Fprintf(w, "%02d:%02d", minutes, seconds)
}

func StartTimer(w http.ResponseWriter, r *http.Request) {
	sessions := middleware.GetSession(w, r)
	start_time := time.Now().Unix()
	workDur, ok := sessions.Values["work_duration"].(int)
	if !ok {
		workDur = 1
	}
	duration := workDur * 60
	is_running := true

	sessions.Values["start_time"] = start_time
	sessions.Values["duration"] = duration
	sessions.Values["is_running"] = is_running
	sessions.Save(r, w)
	log.Println("saved session:", sessions.Values)

	fmt.Fprint(w, "<h1>time started</h1>")

}

func StopTimer(w http.ResponseWriter, r *http.Request) {
	sessions := middleware.GetSession(w, r)
	is_running := false

	delete(sessions.Values, "start_time")
	delete(sessions.Values, "duration")
	sessions.Values["is_running"] = is_running
	sessions.Save(r, w)

	fmt.Fprint(w, "<h1>stopped timer</h1>")

}
