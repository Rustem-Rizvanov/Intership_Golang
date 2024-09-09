package service

import (
	"testing"
	"time"
	"Golang_Intership/internal/domain"
)

type MockUserDB struct{}

func (db *MockUserDB) CreateUser(user *domain.User) error {
	return nil
}

func (db *MockUserDB) GetUserByTelegramID(telegramID int64) (*domain.User, error) {
	return &domain.User{TelegramID: telegramID}, nil
}

func TestCheckAndUpdateUserRequests(t *testing.T) {
	db := &MockUserDB{}
	userService := NewUserService(db) 

	currentTime := time.Now()
	user := &domain.User{TelegramID: 12345, Requests: 5, LastReset: currentTime.Add(-30 * time.Minute)}
	db.CreateUser(user)

	canProceed, err := userService.CheckAndUpdateUserRequests(12345, currentTime)
	if err != nil {
		t.Fatalf("Error checking and updating user requests: %v", err)
	}
	if !canProceed {
		t.Errorf("Expected canProceed to be true, got false")
	}
}
