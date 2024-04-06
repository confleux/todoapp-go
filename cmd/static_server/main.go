package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"web-confleux/internal/config"
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

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil); err != nil {
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
