package service

import (
    "time"
    "Golang_Intership/internal/domain"
)

type UserService struct {
    userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) CheckAndUpdateUserRequests(telegramID int64, currentTime time.Time) (bool, error) {
    user, err := s.userRepo.GetUserByTelegramID(telegramID)
    if err != nil {
        return false, err
    }

    // Если пользователь не найден, просим репозиторий создать его
    if user == nil {
        user = &domain.User{
            TelegramID: telegramID,
            Requests:   1,
            LastReset:  currentTime,
        }
        err := s.userRepo.CreateUser(user)
        if err != nil {
            return false, err
        }
        return true, nil
    }

    if currentTime.Sub(user.LastReset) > time.Hour {
        s.userRepo.ResetUserRequests(user)
        return true, nil
    }

    if user.Requests >= 10 {
        return false, nil
    }

    user.Requests++
    err = s.userRepo.UpdateUserRequests(user)
    if err != nil {
        return false, err
    }

    return true, nil
}
