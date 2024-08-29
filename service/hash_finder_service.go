package service

import "Golang_Intership/domain"

// HashCrackerService обрабатывает запросы на перебор хэшей
type HashCrackerService struct {
	hashService domain.HashService
}

// NewHashCrackerService создает новый экземпляр HashCrackerService
func NewHashCrackerService(hashService domain.HashService) *HashCrackerService {
	return &HashCrackerService{
		hashService: hashService,
	}
}

// CrackHash выполняет перебор и возвращает исходный текст
func (s *HashCrackerService) CrackHash(hash string) (string, error) {
	return s.hashService.FindOriginalText(hash)
}
