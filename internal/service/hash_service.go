package service

import (
    "crypto/md5"
    "encoding/hex"
)

type HashService struct{}

func NewHashService() *HashService {
    return &HashService{}
}

func (s *HashService) HashMessage(message string) (string, error) {
    hash := md5.New()
    hash.Write([]byte(message))
    hashed := hex.EncodeToString(hash.Sum(nil))
    return hashed, nil
}
