package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/alinux78/ulrshortener/internal/service"
)

type uRLShortener struct {
	service service.URLShortener
}

type urlShortenRequest struct {
	URL string `json:"url"`
}

type urlShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func NewURLShortener(svc service.URLShortener) *uRLShortener {
	return &uRLShortener{service: svc}
}

func (h *uRLShortener) Shorten(w http.ResponseWriter, r *http.Request) {
	slog.Debug("shorten request")
	var req urlShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.Shorten(req.URL)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := urlShortenResponse{ShortURL: shortURL}
	json.NewEncoder(w).Encode(resp)
}

func (h *uRLShortener) Resolve(w http.ResponseWriter, r *http.Request) {
	//send not implemented
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
