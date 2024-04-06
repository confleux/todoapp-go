package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment string `env:"ENVIRONMENT" env-description:"app environment"`
	Host        string `env:"HOST" env-description:"http server host" env-required:"true"`
	Port        string `env:"PORT" env-description:"http server port" env-required:"true"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Unable to load .env file: %s", err)
	}

	return &cfg
}
