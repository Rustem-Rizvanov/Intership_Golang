package service

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "Golang_Intership/internal/mocks"
)

// MockUserRepository - мок для репозитория пользователя
type MockUserRepository struct {
    mock.Mock
}

// Пример метода, который может понадобиться в тестах
func (m *MockUserRepository) Create(user domain.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func TestHashService(t *testing.T) {
    mockRepo := new(mocks.MockUserRepository)
    hashService := NewHashService() // Замените на реальный конструктор

    // Пример вызова метода
    hashedPassword, err := hashService.HashPassword("password")
    assert.NoError(t, err)
    assert.NotEmpty(t, hashedPassword)
}
