package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/joho/godotenv"
    "Golang_Intership/config"
    "Golang_Intership/internal/service"
    "Golang_Intership/internal/telegram"
)

func main() {
    // Загружаем переменные окружения из .env файла
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    cfg := config.LoadConfig()

    hashService := service.NewHashService()
    controller := telegram.NewController(hashService)
    bot := telegram.NewBot(cfg.TelegramToken, controller)

    // Контекст для graceful shutdown
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go func() {
        stop := make(chan os.Signal, 1)
        signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
        <-stop
        log.Println("Shutting down gracefully...")
        cancel()
    }()

    go bot.Start()

    <-ctx.Done()

    log.Println("Bot has stopped")
    time.Sleep(1 * time.Second) // ждем завершения всех горутин
}
