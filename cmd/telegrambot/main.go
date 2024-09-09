package main

import (
    "context"
    "database/sql"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
    "Golang_Intership/internal/service"
    "Golang_Intership/internal/telegram"
    "Golang_Intership/config"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
    // Загрузка конфигураций из .env файла
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Загрузка конфигураций
    cfg := config.LoadConfig()

    // Подключение к базе данных
    db, err := sql.Open("postgres", cfg.PostgresURL)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    defer db.Close()


    // Инициализация сервисов
    hashService := service.NewHashService()
    userService := service.NewUserService(db)

    // Инициализация контроллера
    controller := telegram.NewController(userService, hashService)

    // Инициализация бота
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
    time.Sleep(1 * time.Second) // Ждем завершения всех горутин
}
