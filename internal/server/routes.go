package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/kushal0926/pomodoro/internal/handlers"
	"github.com/kushal0926/pomodoro/internal/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/assets"))))

	// registering routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("web/templates/index.html")
		if err != nil {
			log.Fatalf("failed to load template: %v", err)
		}
		session := middleware.GetSession(w, r)
		tmpl.Execute(w, nil)
		log.Println(session)

	})

	mux.HandleFunc("/timer", handlers.GetTimer)
	mux.HandleFunc("/timer/start", handlers.StartTimer)
	mux.HandleFunc("/timer/stop", handlers.StopTimer)
	mux.HandleFunc("/timer/settings", handlers.SaveSettings)

	return s.corsMiddleware(mux)
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}
