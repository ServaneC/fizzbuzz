package main

import (
	"fizzbuzz/pkg/handler"
	"fizzbuzz/pkg/stats"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	statsService := stats.NewStats()
	h := &handler.Handler{
		Stats: statsService,
	}

	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz", h.FizzBuzzHandler).Methods(http.MethodPost)
	r.HandleFunc("/stats", h.StatsHandler).Methods(http.MethodGet)

	fmt.Println("Server running on http://localhost:8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		fmt.Println(err)
	}
}
