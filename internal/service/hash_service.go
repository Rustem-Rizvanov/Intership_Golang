package service

import (
	"crypto/md5"
	"encoding/hex"
	"Golang_Intership/internal/domain"
)

type HashServiceImpl struct{}

func NewHashService() domain.HashService {
	return &HashServiceImpl{}
}

func (s *HashServiceImpl) FindMD5Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
