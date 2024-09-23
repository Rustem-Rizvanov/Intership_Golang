package repository

import (
	"Golang_Intership/internal/domain"

	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUserByTelegramID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database connection: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT id, telegram_id, requests, last_reset FROM users WHERE telegram_id = \$1`).
		WithArgs(12345).
		WillReturnRows(sqlmock.NewRows([]string{"id", "telegram_id", "requests", "last_reset"}).
			AddRow(1, 12345, 5, time.Now()))

	repo := NewUserRepository(db)
	user, err := repo.GetUserByTelegramID(12345)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if user == nil {
		t.Fatal("Expected user, got nil")
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database connection: %v", err)
	}
	defer db.Close()

	mock.ExpectExec(`INSERT INTO users \(telegram_id, requests, last_reset\) VALUES \(\$1, \$2, \$3\)`).
		WithArgs(12345, 5, time.Now()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewUserRepository(db)
	err = repo.CreateUser(&domain.User{TelegramID: 12345, Requests: 5, LastReset: time.Now()})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
