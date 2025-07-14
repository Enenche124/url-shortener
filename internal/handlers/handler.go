package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Enenche124/url-shortener/internal/models"
	"github.com/Enenche124/url-shortener/internal/service"
)

type Handler struct {
	Service *service.URLService
}

func (h *Handler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.URLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	shortCode := h.Service.ShortenURL(req.URL)

	resp := models.URLResponse{
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortCode),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := strings.TrimPrefix(r.URL.Path, "/")
	originalURL, found := h.Service.GetOriginalURL(shortCode)

	if !found {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func MakeShortenHandler(svc *service.URLService) http.HandlerFunc {
	h := &Handler{Service: svc}
	return h.ShortenURLHandler
}

func MakeRedirectHandler(svc *service.URLService) http.HandlerFunc {
	h := &Handler{Service: svc}
	return h.RedirectHandler
}
