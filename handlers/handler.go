package handlers

import (
	"fmt"
	"net/http"
	"url-shortener/models"
	"url-shortener/storage"
)

type Handler struct {
	store *storage.Store
}

func NewHandler(s *storage.Store) *Handler {
	return &Handler{store: s}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		panic("Unable to pass the request body")
	}
	longURL := r.FormValue("url")

	if longURL == "" {
		http.Error(w, "missing 'url' parameter", http.StatusBadRequest)
		return
	}

	shortURL := models.GenerateShortURL()
	h.store.Save(shortURL, longURL)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("https://shorten.url/" + shortURL))
}

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}

	shortURL := r.URL.Path[1:]
	fmt.Println("ShortURL: ", shortURL)

	longURL, exists := h.store.Get(shortURL)
	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}
