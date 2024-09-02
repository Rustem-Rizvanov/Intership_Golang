package telegram

import (
	"Golang_Intership/internal/domain"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Controller struct {
	service domain.HashService
}

func NewController(service domain.HashService) *Controller {
	return &Controller{service: service}
}

func (c *Controller) HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil { // If we got a message
		md5Hash := c.service.FindMD5Hash(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, md5Hash)
		bot.Send(msg)
	}
}
