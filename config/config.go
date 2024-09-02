package config

import (
	"log"
	"os"
)

type Config struct {
	TelegramToken string
}

func LoadConfig() *Config {
	token, exists := os.LookupEnv("TELEGRAM_TOKEN")
	if !exists {
		log.Fatal("TELEGRAM_TOKEN not set in environment variables")
	}

	return &Config{
		TelegramToken: token,
	}
}
