package http

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/zyxevls/internal/usecase"
)

type Handler struct {
	usecase *usecase.URLUseCase
}

func NewHandler(u *usecase.URLUseCase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL         string `json:"url"`
		CustomAlias string `json:"custom_alias,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.usecase.CreateShortURL(req.URL, req.CustomAlias, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"short_url": os.Getenv("BASE_URL") + result.ShortCode,
	})
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")

	original, err := h.usecase.GetOriginalURL(code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	go h.usecase.IncrementClick(code)

	http.Redirect(w, r, original, http.StatusFound)
}
