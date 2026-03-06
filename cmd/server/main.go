package main

import (
	"fizzbuzz/configs"
	"fizzbuzz/internal/handlers"
	"fizzbuzz/internal/stats"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg := configs.Load()
	statsService := stats.NewStats()
	h := &handlers.Handler{
		Stats: statsService,
	}

	r := mux.NewRouter()
	r.Use(LoggingMiddleware)
	r.HandleFunc("/fizzbuzz", h.FizzBuzzHandler).Methods(http.MethodPost)
	r.HandleFunc("/stats", h.StatsHandler).Methods(http.MethodGet)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./internal/static/")))

	log.Printf("Server running on :%s\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Println(err)
	}
}
