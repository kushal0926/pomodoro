package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kushal0926/pomodoro/internal/middleware"
)

func SaveSettings(w http.ResponseWriter, r *http.Request) {
	sessions := middleware.GetSession(w, r)

	r.ParseForm()

	work_duration := r.FormValue("work_duration")

	workDur, err := strconv.Atoi(work_duration)
	if err != nil {
		workDur = 1
	}

	break_duration := r.FormValue("break_duration")
	breakDur, err := strconv.Atoi(break_duration)
	if err != nil {
		breakDur = 1
	}

	sessions.Values["work_duration"] = workDur
	sessions.Values["break_duration"] = breakDur
	sessions.Save(r, w)

	fmt.Fprint(w, "<h1>35/10</h1>")

}
