package handlers

import (
	"encoding/json"
	"errors"
	"fizzbuzz/internal/fizzbuzz"
	"fizzbuzz/internal/stats"
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
		WriteError(w, http.StatusBadRequest, ErrInvalidJSON, err)
		return
	}

	if err := params.IsValid(); err != nil {
		WriteError(w, http.StatusBadRequest, ErrInvalidParameters, err)
		return
	}

	h.Stats.Increment(params)

	result := fizzbuzz.Generate(params)

	WriteJSON(w, http.StatusOK, FizzBuzzResponse{
		Result: result,
	})
}

// GET /stats
func (h *Handler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	params, hits, found := h.Stats.GetMostFrequent()
	if !found {
		WriteError(w, http.StatusNotFound, ErrNotFound, errors.New("no data yet"))
		return
	}

	WriteJSON(w, http.StatusOK, StatsResponse{
		Params: params,
		Hits:   hits,
	})
}
