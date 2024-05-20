package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"

	"client-service/internal/config"
	"client-service/internal/controller"
	"client-service/internal/repository"
	auth "client-service/internal/service/auth"
)

const (
	environmentProd = "prod"
	environmentDev  = "dev"
)

func main() {
	// Basic init
	cfg := config.MustLoad()

	// TODO: Do I actually need logger?
	log := initLogger(cfg.Environment)

	log.Debug("Debug logs are enabled")
	log.Info("Starting app...", slog.String("environment", cfg.Environment))

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve public
	r.Get("/", controller.IndexHandler)
	r.Get("/form", controller.FormHandler)
	r.Get("/todo", controller.TodoHandler)
	r.Get("/login", controller.LoginHandler)
	r.Get("/signup", controller.SignupHandler)

	fs := http.FileServer(http.Dir("./public/src"))
	r.Handle("/src/*", http.StripPrefix("/src/", fs))

	// Create db conn pool
	pool, err := pgxpool.New(context.Background(), cfg.Postgres.Url)
	if err != nil {
		log.Error("unable to init connection: %w", err)
	}
	defer pool.Close()

	// Create repo
	userRepo := repository.NewUserRepository(pool)

	// Create service
	authService, _ := auth.NewAuthService(cfg.Firebase.ServiceAccountConfigPath)

	// Create controller
	authController := controller.NewAuthController(authService, userRepo)

	// API endpoints
	r.Post("/api/signup", authController.SignUp)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTPServer.Port), r); err != nil {
		log.Error("Unable to start server")
	}
}

func initLogger(environment string) *slog.Logger {
	var log *slog.Logger

	switch environment {
	case environmentProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case environmentDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return log
}
