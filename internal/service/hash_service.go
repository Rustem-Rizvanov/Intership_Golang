package service

import (
    "crypto/md5"
    "encoding/hex"
    "errors"
    "sync"
)

type HashRepository interface {
    Save(word string, hash string) error
}

type HashService struct {
    repo HashRepository
}

func NewHashService(repo HashRepository) *HashService {
    return &HashService{repo: repo}
}

func (s *HashService) HashMessage(message string) (string, error) {
    hash := md5.New()
    _, err := hash.Write([]byte(message))
    if err != nil {
        return "", err
    }
    return hex.EncodeToString(hash.Sum(nil)), nil
}

func (s *HashService) SaveWordHash(word string) error {
    hash, err := s.HashMessage(word)
    if err != nil {
        return err
    }
    return s.repo.Save(word, hash)
}

func (s *HashService) BruteForceMD5(md5Hash string, maxLength int, charSet string) (string, error) {
    numWorkers := 10 
    results := make(chan string, 1)
    var wg sync.WaitGroup

    found := false
    var foundMutex sync.Mutex

    worker := func(subset []string, length int) {
        defer wg.Done()

        for _, attempt := range subset {
            if found { 
                return
            }

            hash := md5.New()
            hash.Write([]byte(attempt))
            candidateHash := hex.EncodeToString(hash.Sum(nil))

            if candidateHash == md5Hash {
                foundMutex.Lock()
                if !found {
                    results <- attempt
                    found = true
                }
                foundMutex.Unlock()
                return
            }
        }
    }

    for length := 1; length <= maxLength; length++ {
        combinations := generateCombinations(charSet, "", length)

        subsetSize := len(combinations) / numWorkers
        for i := 0; i < numWorkers; i++ {
            start := i * subsetSize
            end := start + subsetSize
            if i == numWorkers-1 {
                end = len(combinations)
            }

            wg.Add(1)
            go worker(combinations[start:end], length)
        }
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    result, ok := <-results
    if ok {
        return result, nil
    }

    return "", errors.New("не удалось взломать MD5 хэш")
}

func generateCombinations(charSet string, prefix string, length int) []string {
    if length == 0 {
        return []string{prefix}
    }

    var results []string
    for _, char := range charSet {
        newPrefix := prefix + string(char)
        results = append(results, generateCombinations(charSet, newPrefix, length-1)...)
    }

    return results
}

func (s *HashService) HandleMessage(message string) (string, error) {
    if isMD5Hash(message) {
        crackedPassword, err := s.BruteForceMD5(message, 4, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
        if err != nil {
            return "", err
        }
        return "Взломанный пароль: " + crackedPassword, nil
    } else {
        hashedMessage, err := s.HashMessage(message)
        if err != nil {
            return "", err
        }

        if saveErr := s.SaveWordHash(message); saveErr != nil {
            return "", saveErr
        }

        return "Хэшированное сообщение: " + hashedMessage, nil
    }
}


func isMD5Hash(s string) bool {
    if len(s) != 32 {
        return false
    }
    for _, c := range s {
        if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
            return false
        }
    }
    return true
}