package handler

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Nikita-Astafyev/url-shortener/internal/storage"
)

type URLHandler struct {
	storage storage.URLStorage
}

func NewURLHandler(storage storage.URLStorage) *URLHandler {
	return &URLHandler{storage: storage}
}

func isValidURL(rawUrl string) bool {
	_, err := url.ParseRequestURI(rawUrl)
	return err == nil
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if !isValidURL(url) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	originalURL := r.FormValue("url")
	shortURL, err := h.storage.CreateURL(originalURL)
	if err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(shortURL))
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/r/")
	originalURL, err := h.storage.GetOriginalURL(shortURL)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "URL not found", http.StatusNotFound)
		} else {
			log.Printf("DB error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	go func() {
		if err := h.storage.IncrementVisits(shortURL); err != nil {
			log.Printf("Failed to increment visits: %v", err)
		}
	}()

	http.Redirect(w, r, originalURL, http.StatusFound)
}
