package config

import (
	"log"
	"os"
)

type Config struct {
	TelegramToken string
	PostgresURL   string
	RabbitMQURL   string
}

func LoadConfig() *Config {
	return &Config{
		TelegramToken: getEnv("TELEGRAM_TOKEN", ""),
		PostgresURL:   getEnv("POSTGRES_URL", ""),
		RabbitMQURL:   getEnv("RABBITMQ_URL", "amqp://guest:guest@rabbitmq_service:5672/"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		if defaultValue == "" {
			log.Fatalf("Environment variable %s not set", key)
		}
		return defaultValue
	}
	return value
}
