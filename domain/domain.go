package domain

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

// HashService представляет интерфейс для работы с хэшами
type HashService interface {
	FindOriginalText(hash string) (string, error)
}

// SimpleHashService - простая реализация HashService для демонстрации
type SimpleHashService struct{}

// ищет исходный текст по хэшу
func (s SimpleHashService) FindOriginalText(hash string) (string, error) {
	dictionary := map[string]string{
		"secret": "5ebe2294ecd0e0f08eab7690d2a6ee69", // md5("secret")
	}

	for text, h := range dictionary {
		if h == hash {
			return text, nil
		}
	}
	return "", errors.New("original text not found")
}

// HashText возвращает md5-хэш от текста
func HashText(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
