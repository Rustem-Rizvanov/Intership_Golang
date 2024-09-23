package service

import (
    "encoding/json"
    "log"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "Golang_Intership/internal/domain"
)

type MessageService struct {
    broker     domain.MessageBroker
    bot        *tgbotapi.BotAPI
    hashService domain.HashService 
}

func NewMessageService(broker domain.MessageBroker, bot *tgbotapi.BotAPI, hashService domain.HashService) *MessageService {
    return &MessageService{broker: broker, bot: bot, hashService: hashService}
}

func (s *MessageService) Start() error {
    go func() {
        err := s.ReceiveMessage("telegram_queue", s.handleMessage)
        if err != nil {
            log.Fatalf("Failed to consume messages: %v", err)
        }
    }()
    return nil
}

func (s *MessageService) SendMessage(telegramID int64, message string) error {
    telegramMsg := tgbotapi.NewMessage(telegramID, message)
    _, err := s.bot.Send(telegramMsg)
    return err
}

func (s *MessageService) ReceiveMessage(queue string, handler func([]byte) error) error {
    return s.broker.Consume(queue, handler)
}

func (s *MessageService) handleMessage(msg []byte) error {
    log.Printf("Received a message: %s", msg)

    var request struct {
        TelegramID int64  `json:"telegram_id"`
        Message    string `json:"message"`
    }

    if err := json.Unmarshal(msg, &request); err != nil {
        log.Printf("Failed to unmarshal message: %v", err)
        return err
    }

    response, err := s.hashService.HandleMessage(request.Message)
    if err != nil {
        return err
    }

    return s.SendMessage(request.TelegramID, response)
}

func (s *MessageService) HandleUserMessage(telegramID int64, message string) error {
    msg := struct {
        TelegramID int64  `json:"telegram_id"`
        Message    string `json:"message"`
    }{
        TelegramID: telegramID,
        Message:    message,
    }

    data, err := json.Marshal(msg)
    if err != nil {
        return err
    }

    return s.broker.Publish("telegram_queue", data)
}
