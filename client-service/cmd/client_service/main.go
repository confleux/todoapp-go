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

	_ "client-service/docs"
	"client-service/internal/config"
	"client-service/internal/controller"
	middleware2 "client-service/internal/middleware"
	"client-service/internal/repository"
	"client-service/internal/routes"
	"client-service/internal/service"
	auth "client-service/internal/service/auth"
)

const (
	environmentProd = "prod"
	environmentDev  = "dev"
)

// @title           Todo API
// @version         1.0
//// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/

//// @securityDefinitions.basic  BasicAuth

// // @externalDocs.description  OpenAPI
// // @externalDocs.url          https://swagger.io/resources/open-api/
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// programmatically set swagger info
	//docs.SwaggerInfo.Title = "Todo API"
	//docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "web-confleux.onrender.com"
	//docs.SwaggerInfo.BasePath = "/api"
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}

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

	// Create service
	authService, _ := auth.NewAuthService(cfg.Firebase.ServiceAccountConfigPath, userRepo)
	todoService := service.NewTodoService(todoRepo)

	// Create controller
	authController := controller.NewAuthController(authService)
	todoController := controller.NewTodoController(todoService)

	// API endpoints
	publicApiGroup := r.Group(nil)
	publicApiGroup.Post("/api/signup", authController.SignUp)

	privateApiGroup := r.Group(nil)
	privateApiGroup.Use(middleware2.Auth(authService))
	privateApiGroup.Mount("/api/todos", routes.NewTodoResource(todoController).Routes())

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), //The url pointing to API definition
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
