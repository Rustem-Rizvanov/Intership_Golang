package telegram

import (
	"testing"
	"Golang_Intership/internal/service"
)

type MockController struct{}

func (c *MockController) HandleMessage(user *service.User, message string) (string, error) {
	return "mock response", nil
}

func TestNewBot(t *testing.T) {
	mockController := &MockController{}
	token := "test-token"

	bot := NewBot(token, mockController)
	if bot.API == nil {
		t.Fatal("Expected bot.API to be initialized")
	}
	if bot.Controller != mockController {
		t.Errorf("Expected Controller to be %v, got %v", mock
