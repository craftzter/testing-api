package routes

import (
	"monly-login-api/internal/handlers"
	"monly-login-api/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupUserRoute(r *chi.Mux, handler *handlers.Handler) {
	r.Route("/users", func(r chi.Router) {
		// PUBLIC ROUTES
		r.Post("/register", handler.CreateUserHandler())
		r.Post("/login", handler.LoginHandler())
		r.Get("/{id}", handler.GetUserByIDHandler())
		r.Get("/list", handler.GetListUser())

		// PROTECTED ROUTES
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Put("/profile/{id}", handler.UpdateHandler())
		})
	})
}
