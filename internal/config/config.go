package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Environment string `yaml:"environment" envDefault:"development"`
	HTTPServer  HTTPServerConfig
	Postgres    PostgresServerConfig
	Firebase    FirebaseConfig
}

type HTTPServerConfig struct {
	Port int `yaml:"port"`
}

type PostgresServerConfig struct {
	Url string `yaml:"url"`
}

type FirebaseConfig struct {
	ServiceAccountConfigPath string `yaml:"service_account_config_path"`
}

func MustLoad() *Config {
	configPath := getConfigPath()
	if configPath == "" {
		log.Fatalln("Config file path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s ", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Unable to load config file: %v", err)
	}

	return &cfg
}

func getConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "config file path")
	flag.Parse()

	return res
}
