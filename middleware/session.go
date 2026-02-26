package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func GetSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return session
}
