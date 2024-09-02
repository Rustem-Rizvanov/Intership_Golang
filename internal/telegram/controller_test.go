package telegram

import (
	"testing"

	"github.com/golang/mock/gomock"
	"Golang_Intership/internal/domain"
	"Golang_Intership/internal/mocks"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func TestController_HandleMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockHashService(ctrl)
	controller := NewController(mockService)

	mockService.EXPECT().FindMD5Hash("hello").Return("5d41402abc4b2a76b9719d911017c592")

	bot := &tgbotapi.BotAPI{}
	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "hello",
			Chat: &tgbotapi.Chat{ID: 1234},
		},
	}

	controller.HandleMessage(bot, update)
}
