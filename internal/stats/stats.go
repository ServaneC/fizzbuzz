package stats

import (
	"fizzbuzz/internal/fizzbuzz"
	"sync"
)

type Stats struct {
	mu        sync.Mutex
	store     map[fizzbuzz.Params]int
	maxParams fizzbuzz.Params
	maxHits   int
}

func NewStats() *Stats {
	return &Stats{
		store: make(map[fizzbuzz.Params]int),
	}
}

func (s *Stats) Increment(p fizzbuzz.Params) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[p]++

	count := s.store[p]
	if count > s.maxHits {
		s.maxHits = count
		s.maxParams = p
	}
}

func (s *Stats) GetMostFrequent() (fizzbuzz.Params, int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.maxHits == 0 {
		return fizzbuzz.Params{}, 0, false
	}

	return s.maxParams, s.maxHits, true
}
