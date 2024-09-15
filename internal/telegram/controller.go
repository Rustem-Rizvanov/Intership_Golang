package telegram

import (
    "Golang_Intership/internal/domain"
    "Golang_Intership/internal/service"
    "time"

)

type Controller struct {
    userService *service.UserService
    hashService *service.HashService
}

func NewController(userService *service.UserService, hashService *service.HashService) *Controller {
    return &Controller{
        userService: userService,
        hashService: hashService,
    }
}

func (c *Controller) HandleMessage(user *domain.User, message string) (string, error) {
    currentTime := time.Now()

    canProceed, err := c.userService.CheckAndUpdateUserRequests(user.TelegramID, currentTime)
    if err != nil {
        return "", err
    }
    if !canProceed {
        resetTime := currentTime.Add(time.Second).Format("15:04:05")
        return "Ты в камере, жди еще: " + resetTime, nil
    }

    var response string
    if isMD5Hash(message) {
        // Если хэш то взлом
        crackedPassword, err := c.hashService.BruteForceMD5(message, 4, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
        if err != nil {
            return "", err
        }
        response = "Взломанный пароль: " + crackedPassword
    } else {
        hashedMessage, err := c.hashService.HashMessage(message)
        if err != nil {
            return "", err
        }
        response = "Хэшированное сообщение: " + hashedMessage
    }

    return response, nil
}

// является ли сообщение хэшем
func isMD5Hash(s string) bool {
    if len(s) != 32 {
        return false
    }
    // олько допустимые символы 
    for _, c := range s {
        if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
            return false
        }
    }
    return true
}
