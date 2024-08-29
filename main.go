package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"Golang_Intership/domain"
	"Golang_Intership/service"
)
func main() {
	// Инициализация Telegram бота
	bot, err := tgbotapi.NewBotAPI("7259501029:AAE1XtlPeUwFjbX1C73QkVkBywEqoygZewQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// Создание сервиса и доменной логики
	hashService := domain.SimpleHashService{}
	hashFinderService := service.NewHashFinderService(hashService)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Обработка сообщений
	for update := range updates {
		if update.Message == nil {
			continue
		}

		md5Hash := update.Message.Text
		originalText, err := hashFinderService.FindAndReturnOriginal(md5Hash)

		var response string
		if err != nil {
			response = "Original text not found for this hash"
		} else {
			response = "Original text: " + originalText
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		bot.Send(msg)
	}
}