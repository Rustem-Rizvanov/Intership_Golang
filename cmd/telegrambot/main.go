package main

import (
    "context"
    "database/sql"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "Golang_Intership/internal/service"
    "Golang_Intership/internal/telegram"
    "Golang_Intership/internal/statistics"
    "Golang_Intership/config"
    "Golang_Intership/internal/repository"
    "Golang_Intership/internal/adapter"
    "github.com/joho/godotenv"
    "github.com/go-telegram-bot-api/telegram-bot-api"
    _ "github.com/lib/pq"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    cfg := config.LoadConfig()
    log.Println("Configurations loaded successfully.")

    db, err := sql.Open("postgres", cfg.PostgresURL)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    defer db.Close()
    log.Println("Connected to the database successfully.")

    loggingDB, err := sql.Open("postgres", os.Getenv("LOGGING_POSTGRES_URL"))
    if err != nil {
        log.Fatalf("Error connecting to logging database: %v", err)
    }
    defer loggingDB.Close()
    log.Println("Connected to the Second database successfully")

    userRepo := repository.NewUserRepository(db)
    hashRepo, err := repository.NewPostgresHashRepository(os.Getenv("HASH_POSTGRES_URL"))
    if err != nil {
        log.Fatalf("Error initializing hash repository: %v", err)
    }
    log.Println("Repositories initialized.")

    hashService := service.NewHashService(hashRepo)
    userService := service.NewUserService(userRepo)
    log.Println("Services initialized.")

    rabbitMQAdapter, err := adapter.NewRabbitMQAdapter(cfg.RabbitMQURL)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer rabbitMQAdapter.Close()
    log.Println("RabbitMQ adapter initialized.")

    bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
    if err != nil {
        log.Fatalf("Failed to create Telegram bot: %v", err)
    }
    log.Println("Telegram bot initialized.")

    messageService := service.NewMessageService(rabbitMQAdapter, bot, hashService)
    log.Println("Message service initialized.")

    controller := telegram.NewController(userService, hashService, messageService, bot, rabbitMQAdapter)
    log.Println("Controller initialized.")

    // Сервис статистики
    statsRepo := statistics.NewStatisticsRepository(loggingDB)
    statsService := statistics.NewStatisticsService(statsRepo)
    statsController := statistics.NewStatisticsController(statsService)

    // HTTP-сервер для статистики
    http.HandleFunc("/statistics", statsController.GetStatisticsHandler)
    go func() {
        log.Println("Starting statistics HTTP server on :8081")
        if err := http.ListenAndServe(":8081", nil); err != nil {
            log.Fatalf("Failed to start statistics HTTP server: %v", err)
        }
    }()

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go func() {
        stop := make(chan os.Signal, 1)
        signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
        <-stop
        log.Println("Shutting down gracefully...")
        cancel()
    }()

    go func() {
        if err := messageService.Start(); err != nil {
            log.Fatalf("Error starting message service: %v", err)
        }
    }()

    go func() {
        if err := controller.Start(); err != nil {
            log.Fatalf("Error starting controller: %v", err)
        }
    }()

    <-ctx.Done()
    log.Println("Services have stopped.")
    time.Sleep(1 * time.Second)
}
