package routes

import (
	"monly-login-api/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupHealthRoute(r *chi.Mux, handler *handlers.Handler) {
	r.Get("/health", handler.HealtHandler())
}
