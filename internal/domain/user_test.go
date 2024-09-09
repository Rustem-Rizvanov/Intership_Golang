package domain

import (
    "testing"
    "time"
)

func TestUserCreation(t *testing.T) {
    user := User{
        ID:         1,
        TelegramID: 12345,
        Requests:   5,
        LastReset:  time.Now(),
    }

    if user.ID != 1 {
        t.Errorf("Expected ID 1, got %d", user.ID)
    }
    if user.TelegramID != 12345 {
        t.Errorf("Expected TelegramID 12345, got %d", user.TelegramID)
    }
}
