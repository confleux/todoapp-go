package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host string `env:"HOST" env-required:"true"`
	Port string `env:"PORT" env-required:"true"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	return &Config{
		Host: host,
		Port: port,
	}
}
