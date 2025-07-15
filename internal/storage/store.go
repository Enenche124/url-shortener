package storage

import "sync"

type URLStore struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewURLStore() *URLStore {
	return &URLStore{
		store: make(map[string]string),
	}
}

func (s *URLStore) Save(shortCode, originalURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[shortCode] = originalURL
}

func (s *URLStore) Get(shortCode string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, exists := s.store[shortCode]
	return url, exists
}
