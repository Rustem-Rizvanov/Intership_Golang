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

        user := &domain.User{TelegramID: int64(update.Message.From.ID)}
        response, err := b.Controller.HandleMessage(user, update.Message.Text)
        if err != nil {
            log.Println("Error handling message:", err)
            continue
        }

        msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
        if _, err := b.API.Send(msg); err != nil {
            log.Println("Failed to send message:", err)
        }
    }
}
