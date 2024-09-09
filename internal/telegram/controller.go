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

    // Проверяем и обновляем количество запросов
    canProceed, err := c.userService.CheckAndUpdateUserRequests(user.TelegramID, currentTime)
    if err != nil {
        return "", err
    }
    if !canProceed {
        resetTime := currentTime.Add(time.Second).Format("15:04:05")
        return "Ты в камере, жди еще: " + resetTime, nil
    }

    // Обрабатываем сообщение с хэшированием
    hashedMessage, err := c.hashService.HashMessage(message)
    if err != nil {
        return "", err
    }

    return "Хэшированное сообщение: " + hashedMessage, nil
}
