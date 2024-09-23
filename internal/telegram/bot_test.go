package telegram

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "Golang_Intership/internal/mocks"
    "Golang_Intership/internal/service"
)

// Пример теста для бота
func TestBot(t *testing.T) {
    mockController := new(mocks.MockController)
    bot := NewBot(mockController)

    // Пример теста
    assert.Equal(t, mockController, bot.Controller)
}
