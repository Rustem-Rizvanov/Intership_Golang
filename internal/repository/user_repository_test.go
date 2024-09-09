package repository

import (
	"testing"
	"time"
	"Golang_Intership/internal/domain"
)


func (db *FakeDB) CreateUser(user *domain.User) error {
	return nil
}

func (db *FakeDB) GetUserByTelegramID(telegramID int64) (*domain.User, error) {
	return &domain.User{TelegramID: telegramID}, nil
}

func TestUserRepository(t *testing.T) {
	db := &FakeDB{}
	repo := NewUserRepository(db) // Проверь сигнатуру NewUserRepository

	user := &domain.User{ID: 1, TelegramID: 12345, Requests: 5, LastReset: time.Now()}
	db.CreateUser(user)

	result, err := repo.GetUserByTelegramID(12345)
	if err != nil {
		t.Fatalf("Error getting user: %v", err)
	}
	if result.TelegramID != 12345 {
		t.Errorf("Expected TelegramID 12345, got %d", result.TelegramID)
	}
}
