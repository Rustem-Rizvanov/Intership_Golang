package service

import "testing"

func TestHashMessage(t *testing.T) {
    service := NewHashService()
    hashed, err := service.HashMessage("test")
    if err != nil {
        t.Errorf("Error hashing message: %v", err)
    }
    // Проверить ожидаемый хэш
    if hashed == "" {
        t.Errorf("Expected non-empty hash, got empty string")
    }
}
