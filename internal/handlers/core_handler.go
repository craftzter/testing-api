package handlers

import (
	"database/sql"
	db "monly-login-api/internal/generate"
	"monly-login-api/internal/service"
)

type Handler struct {
	DB          *sql.DB
	Queries     *db.Queries
	UserService *service.UserService
}

func NewHandlers(db *sql.DB, queries *db.Queries, userService *service.UserService) *Handler {
	return &Handler{
		DB:          db,
		Queries:     queries,
		UserService: userService,
	}
}
