// main.go
package main

import (
	"fmt"
	"log"
	"log/slog"
	"monly-login-api/config"
	"monly-login-api/database"
	"monly-login-api/internal/handlers"
	"monly-login-api/internal/routes"
	"monly-login-api/internal/service"
	"net/http"
	"os"

	db "monly-login-api/internal/generate"

	"github.com/go-chi/chi/v5"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	// load config
	cfg, err := config.LoadingConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// connect to db
	conn := database.ConnectDB(cfg.DatabaseUrl)
	defer conn.Close()

	// sqlc queries
	queries := db.New(conn)

	// bikin user service
	userService := service.NewUserService(queries, logger)

	// bikin handler
	handler := handlers.NewHandlers(conn, queries, userService)

	// router
	r := chi.NewRouter()
	routes.SetupUserRoute(r, handler)
	routes.SetupHealthRoute(r, handler)
	// start server
	addr := fmt.Sprintf(":%s", cfg.AppPort)
	fmt.Printf("🚀 Server running on %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
