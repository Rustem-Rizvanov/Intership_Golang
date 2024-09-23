package service

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "Golang_Intership/internal/mocks"
)

// Пример теста
func TestUserService(t *testing.T) {
    mockMD5Service := new(mocks.MockMD5Service)
    userService := NewUserService(mockMD5Service)

    // Пример вызова метода
    hashedPassword, err := userService.HashPassword("password")
    assert.NoError(t, err)
    assert.NotEmpty(t, hashedPassword)
}
