package telegram

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "Golang_Intership/internal/mocks"
    "Golang_Intership/internal/service"
)

// Пример теста для контроллера
func TestUserController(t *testing.T) {
    mockUserService := new(mocks.MockUserService)
    controller := NewUserController(mockUserService)

    // Пример теста
    user := service.User{ID: 1, Name: "Test"}
    mockUserService.On("GetUserByID", 1).Return(user, nil)

    fetchedUser, err := controller.GetUser(1)
    assert.NoError(t, err)
    assert.Equal(t, user, fetchedUser)
}
