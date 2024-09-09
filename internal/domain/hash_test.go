package domain

import "testing"

type MockHashService struct{}

func (s *MockHashService) FindMD5Hash(input string) string {
    return "mocked_hash"
}

func TestFindMD5Hash(t *testing.T) {
    service := &MockHashService{}
    result := service.FindMD5Hash("test")
    expected := "mocked_hash"
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}
