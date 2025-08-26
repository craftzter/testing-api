package routes

import (
	"monly-login-api/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupUserRoute(r *chi.Mux, handler *handlers.Handler) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", handler.CreateUserHandler())
		r.Post("/login", handler.LoginHandler())
		r.Put("/{id}", handler.UpdateHandler())
	})
}
