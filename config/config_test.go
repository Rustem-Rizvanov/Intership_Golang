package config

import (
    "os"
    "testing"
)

func TestLoadConfig(t *testing.T) {
    os.Setenv("POSTGRES_URL", "test_url")
    os.Setenv("TELEGRAM_TOKEN", "test_token")
    
    cfg := LoadConfig()

    if cfg.PostgresURL != "test_url" {
        t.Errorf("Expected POSTGRES_URL to be 'test_url', got %s", cfg.PostgresURL)
    }
    if cfg.TelegramToken != "test_token" {
        t.Errorf("Expected TELEGRAM_TOKEN to be 'test_token', got %s", cfg.TelegramToken)
    }
}
