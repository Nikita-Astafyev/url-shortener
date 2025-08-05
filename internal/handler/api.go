package handler

import (
	"net/http"

	"github.com/Nikita-Astafyev/url-shortener/internal/storage"
)

type URLHandler struct {
	storage storage.URLStorage
}

func NewURLHandler(storage storage.URLStorage) *URLHandler {
	return &URLHandler{storage: storage}
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	originalURL := r.FormValue("url")
	shortURL, err := h.storage.CreateURL(originalURL)
	if err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(shortURL))
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	originalURL, err := h.storage.GetOriginalURL(shortURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	h.storage.IncrementVisits(shortURL)
	http.Redirect(w, r, originalURL, http.StatusFound)
}
