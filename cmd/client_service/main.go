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
	httpSwagger "github.com/swaggo/http-swagger"

	_ "todoapp-go/docs"
	"todoapp-go/internal/config"
	"todoapp-go/internal/controller"
	middleware2 "todoapp-go/internal/middleware"
	"todoapp-go/internal/repository"
	"todoapp-go/internal/routes"
	"todoapp-go/internal/service"
	auth "todoapp-go/internal/service/auth"
)

const (
	environmentProd = "prod"
	environmentDev  = "dev"
)

// @title           Todo API
// @version         1.0

// @host web-confleux.onrender.com
// @BasePath  /api/
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Basic init
	cfg := config.MustLoad()

	// TODO: Do I actually need logger?
	log := initLogger(cfg.Environment)

	log.Debug("Debug logs are enabled")
	log.Info("Starting app...", slog.String("environment", cfg.Environment))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Serve public
	publicGroup := r.Group(nil)
	publicGroup.Get("/", controller.IndexHandler)
	publicGroup.Get("/form", controller.FormHandler)
	publicGroup.Get("/todo", controller.TodoHandler)
	publicGroup.Get("/login", controller.LoginHandler)
	publicGroup.Get("/signup", controller.SignupHandler)
	publicGroup.Get("/todo-app", controller.TodoAppHandler)

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
	todoRepo := repository.NewTodoRepository(pool)
	feedbackRepo := repository.NewFeedbackRepository(pool)

	// Create service
	authService, _ := auth.NewAuthService(cfg.Firebase.ServiceAccountConfigPath, userRepo)
	todoService := service.NewTodoService(todoRepo)
	feedbackService := service.NewFeedbackService(feedbackRepo)

	// Create controller
	authController := controller.NewAuthController(authService)
	todoController := controller.NewTodoController(todoService)

	// API endpoints
	publicApiGroup := r.Group(nil)
	publicApiGroup.Post("/api/signup", authController.SignUp)

	privateApiGroup := r.Group(nil)
	privateApiGroup.Use(middleware2.Auth(authService))
	privateApiGroup.Mount("/api/todos", routes.NewTodoResource(todoController).Routes())

	r.Get("/ws", feedbackService.Handler)

	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("https://web-confleux.onrender.com/swagger/doc.json"), //The url pointing to API definition
	))

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

func hanler(w http.ResponseWriter, r *http.Request) {

}
