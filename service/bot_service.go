package service

import (
	"Golang_Intership/domain"
)
type HashFinderService struct {
	hashService domain.HashService
}

func NewHashFinderService(hashService domain.HashService) *HashFinderService {
	return &HashFinderService{
		hashService: hashService,
	}
}

// FindAndReturnOriginal ищет оригинальный текст по хэшу и возвращает его
func (s *HashFinderService) FindAndReturnOriginal(hash string) (string, error) {
	return s.hashService.FindOriginalText(hash)
}