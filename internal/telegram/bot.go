package telegram

import (
    "Golang_Intership/internal/domain"
    "log"
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
    API        *tgbotapi.BotAPI
    Controller *Controller
}

func NewBot(token string, controller *Controller) *Bot {
    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        log.Panic(err)
    }

    return &Bot{
        API:        bot,
        Controller: controller,
    }
}

func (b *Bot) Start() {
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := b.API.GetUpdatesChan(u)
    if err != nil {
        log.Fatalf("Failed to get updates channel: %v", err)
    }

    for update := range updates {
        if update.Message == nil { // Игнорируем не сообщения
            continue
        }

        user := &domain.User{
            TelegramID: int64(update.Message.From.ID),
            ChatID:     int64(update.Message.Chat.ID),
        }

        existingUser, err := b.Controller.userService.GetUserByTelegramID(user.TelegramID)
        if err != nil {
            log.Println("Error fetching user:", err)
            continue
        }

        if existingUser == nil {
            err := b.Controller.userService.CreateUser(user)
            if err != nil {
                log.Println("Error creating user:", err)
                continue
            }
        } else {
            if existingUser.ChatID != user.ChatID {
                existingUser.ChatID = user.ChatID
                err := b.Controller.userService.UpdateUser(existingUser)
                if err != nil {
                    log.Println("Error updating user:", err)
                }
            }
        }

        message := update.Message.Text

        err = b.Controller.messageService.SendMessage(user.TelegramID, message)
        if err != nil {
            log.Println("Error sending message to RabbitMQ:", err)
            continue
        }
    }
}
