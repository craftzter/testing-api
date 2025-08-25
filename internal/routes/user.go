package routes

import (
	"monly-login-api/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupUserRoute(r *chi.Mux, handler *handlers.Handler) {
	// karena CreateUserHandler return http.HandlerFunc
	r.Post("/users", handler.CreateUserHandler())
}
