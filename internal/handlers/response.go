package handlers

import (
	"encoding/json"
	"fizzbuzz/internal/fizzbuzz"
	"log"
	"net/http"
)

type FizzBuzzResponse struct {
	Result []string `json:"result"`
}

type StatsResponse struct {
	fizzbuzz.Params
	Hits int `json:"hits"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}
