package telegram

import (
	"testing"
)

type MockUserService struct{}

func (s *MockUserService) SomeMethod() error {
	return nil
}

func TestNewController(t *testing.T) {
	mockUserService := &MockUserService{}
	hashService := &MockHashService{} // Предполагается, что есть такой мок

	controller := NewController(mockUserService, hashService)

	if controller.userService != mockUserService {
		t.Errorf("Expected userService to be %v, got %v", mockUserService, controller.userService)
	}
}
