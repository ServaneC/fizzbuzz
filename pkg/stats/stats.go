package stats

import (
	"fizzbuzz/pkg/fizzbuzz"
	"sync"
)

type Stats struct {
	mu    sync.Mutex
	store map[fizzbuzz.Params]int
}

func NewStats() *Stats {
	return &Stats{
		store: make(map[fizzbuzz.Params]int),
	}
}

func (s *Stats) Increment(key fizzbuzz.Params) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key]++
}

func (s *Stats) GetMostFrequent() (fizzbuzz.Params, int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var maxKey fizzbuzz.Params
	maxCount := 0
	found := false

	for k, v := range s.store {
		if !found || v > maxCount {
			maxKey = k
			maxCount = v
			found = true
		}
	}

	return maxKey, maxCount, found
}
