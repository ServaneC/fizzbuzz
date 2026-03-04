package handler

import (
	"encoding/json"
	"fizzbuzz/pkg/fizzbuzz"
	"fizzbuzz/pkg/stats"
	"net/http"
)

type Handler struct {
	Stats *stats.Stats
}

// POST /fizzbuzz
func (h *Handler) FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	var params fizzbuzz.Params

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if !params.IsValid() {
		http.Error(w, "Invalid or missing parameters", http.StatusBadRequest)
		return
	}

	h.Stats.Increment(params)

	result := fizzbuzz.Generate(params)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GET /stats
func (h *Handler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	params, hits, found := h.Stats.GetMostFrequent()
	if !found {
		http.Error(w, "No data yet", http.StatusNotFound)
		return
	}

	response := struct {
		fizzbuzz.Params
		Hits int `json:"hits"`
	}{
		Params: params,
		Hits:   hits,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
