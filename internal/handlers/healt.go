package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) HealtHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"message": "API is healthy",
			"status":  "Server berjalan dengan baik",
		}
		json.NewEncoder(w).Encode(response)
	}
}
