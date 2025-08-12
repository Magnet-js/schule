package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
	Port        string `env:"PORT,default=8080"`
}

func MustLoadConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is needed but not set")
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	return cfg
}