package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Environment string `yaml:"environment" envDefault:"development"`
	HTTP        HTTPServerConfig
}

type HTTPServerConfig struct {
	Port int `yaml:"port"`
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
