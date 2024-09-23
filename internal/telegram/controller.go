package telegram

import (
    "Golang_Intership/internal/domain"
    "Golang_Intership/internal/service"
	"Golang_Intership/internal/adapter"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "log"
	"time"
)

type Controller struct {
    userService     *service.UserService
    hashService     *service.HashService
    messageService  *service.MessageService
    bot             *tgbotapi.BotAPI
    rabbitMQAdapter *adapter.RabbitMQAdapter 
}

func NewController(
    userService *service.UserService,
    hashService *service.HashService,
    messageService *service.MessageService,
    bot *tgbotapi.BotAPI,
    rabbitMQAdapter *adapter.RabbitMQAdapter, 
) *Controller {
    return &Controller{
        userService:     userService,
        hashService:     hashService,
        messageService:  messageService,
        bot:             bot,
        rabbitMQAdapter: rabbitMQAdapter, 
    }
}


func (c *Controller) Start() error {
    updates, err := c.bot.GetUpdatesChan(tgbotapi.UpdateConfig{})
    if err != nil {
        return err
    }

    for update := range updates {
        if update.Message == nil { 
            continue
        }

        user, err := c.userService.GetUserByTelegramID(update.Message.Chat.ID)
        if err != nil {
            log.Printf("Error getting user: %v", err)
            continue
        }

        err = c.messageService.HandleUserMessage(user.TelegramID, update.Message.Text)
        if err != nil {
            log.Printf("Error sending message to RabbitMQ: %v", err)
            continue
        }
    }

    return nil
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
