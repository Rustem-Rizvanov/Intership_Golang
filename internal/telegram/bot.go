package telegram

import (
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
		log.Panic(err)
	}

	for update := range updates {
		b.Controller.HandleMessage(b.API, update)
	}
}
