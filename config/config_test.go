package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Установите переменные окружения для теста
	os.Setenv("TELEGRAM_TOKEN", "test_token")
	os.Setenv("POSTGRES_URL", "test_postgres_url")

	cfg := LoadConfig()

	if cfg.TelegramToken != "test_token" {
		t.Errorf("Expected TelegramToken to be 'test_token', got %s", cfg.TelegramToken)
	}
	if cfg.PostgresURL != "test_postgres_url" {
		t.Errorf("Expected PostgresURL to be 'test_postgres_url', got %s", cfg.PostgresURL)
	}
}
