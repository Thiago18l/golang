package main

import (
	"testing"
	"time"
)

// SpyStore let you do a simulation with a store
type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

// Fetch will return a answer with a small delay
func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

// Cancel will store the call
func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("store dont was alerted about the cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf(`store was alerted to cancel`)
	}
}
