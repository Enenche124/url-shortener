package storage

import "sync"

// URLStore holds the mappings of short codes to original URLs.
type URLStore struct {
	mu    sync.RWMutex
	store map[string]string
}

// NewURLStore creates and returns a new URLStore.
func NewURLStore() *URLStore {
	return &URLStore{
		store: make(map[string]string),
	}
}

// Save saves a short code and its corresponding long URL to the store.
func (s *URLStore) Save(shortCode, originalURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[shortCode] = originalURL
}

// Get retrieves the original URL for a given short code.
// Returns empty string if not found.
func (s *URLStore) Get(shortCode string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, exists := s.store[shortCode]
	return url, exists
}
