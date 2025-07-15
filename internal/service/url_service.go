package service

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/Enenche124/url-shortener/internal/storage"
)

type URLService struct {
	Store *storage.URLStore
}

func (s *URLService) ShortenURL(originalURL string) string {
	hash := sha1.Sum([]byte(originalURL))
	shortCode := hex.EncodeToString(hash[:])[:6]

	s.Store.Save(shortCode, originalURL)
	return shortCode
}

func (s *URLService) GetOriginalURL(shortCode string) (string, bool) {
	return s.Store.Get(shortCode)
}
