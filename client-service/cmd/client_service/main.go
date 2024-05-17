package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"client-service/internal/config"
	"client-service/internal/front"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	environmentProd = "prod"
	environmentDev  = "dev"
)

func main() {
	cfg := config.MustLoad()

	log := initLogger(cfg.Environment)

	log.Debug("Debug logs are enabled")
	log.Info("Starting app...", slog.String("environment", cfg.Environment))

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", front.IndexHandler)
	r.Get("/form", front.FormHandler)
	r.Get("/todo", front.TodoHandler)

	fs := http.FileServer(http.Dir("./public/src"))
	r.Handle("/src/*", http.StripPrefix("/src/", fs))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTP.Port), r); err != nil {
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
