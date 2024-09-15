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
	"Golang_Intership/internal/repository"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" 
)

func main() {
	// Загрузка конфигураций из .env файла
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// конфигураций
	cfg := config.LoadConfig()
	log.Println("Configurations loaded successfully.")

	// бд
	db, err := sql.Open("postgres", cfg.PostgresURL)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()
	log.Println("Connected to the database successfully.")

	// репозиторий
	userRepo := repository.NewUserRepository(db)
	log.Println("User repository initialized.")

	// сервис
	hashService := service.NewHashService()
	userService := service.NewUserService(userRepo)
	log.Println("Services initialized.")

	// контроллер
	controller := telegram.NewController(userService, hashService)
	log.Println("Controller initialized.")

	// бот
	bot := telegram.NewBot(cfg.TelegramToken, controller)
	log.Println("Bot initialized.")

	// graceful shutdown
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
	time.Sleep(1 * time.Second) 
}
