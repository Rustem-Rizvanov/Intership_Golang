package service

import (
	"testing"
)

func TestFindMD5Hash(t *testing.T) {
	service := NewHashService()

	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		{"world", "7d793037a0760186574b0282f2f435e7"},
		{"GoLang", "e59472ab9dc7b5de90c9c4f32e168bc3"},
		{"", "d41d8cd98f00b204e9800998ecf8427e"}, // Пустая строка
	}

	for _, test := range tests {
		result := service.FindMD5Hash(test.input)
		if result != test.expected {
			t.Errorf("Expected %s but got %s", test.expected, result)
		}
	}
}

func TestFindMD5Hash_EmptyString(t *testing.T) {
	service := NewHashService()

	result := service.FindMD5Hash("")
	expected := "d41d8cd98f00b204e9800998ecf8427e"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestFindMD5Hash_SpecialCharacters(t *testing.T) {
	service := NewHashService()

	result := service.FindMD5Hash("!@#$%^&*()")
	expected := "f709dc4048f7b17b9195e4abff508ff2"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
