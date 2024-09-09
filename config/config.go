package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresURL   string
	TelegramToken string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		PostgresURL:   os.Getenv("POSTGRES_URL"),
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
	}
}
