package repository

import (
	"database/sql"
	"log"
	"time"

	"Golang_Intership/internal/domain"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func ConnectPostgres(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

func (r *UserRepository) GetUserByTelegramID(telegramID int64) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, telegram_id, requests, last_reset FROM users WHERE telegram_id = $1`
	err := r.db.QueryRow(query, telegramID).Scan(&user.ID, &user.TelegramID, &user.Requests, &user.LastReset)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUserRequests(user *domain.User) error {
	query := `UPDATE users SET requests = $1, last_reset = $2 WHERE id = $3`
	_, err := r.db.Exec(query, user.Requests, user.LastReset, user.ID)
	return err
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	query := `INSERT INTO users (telegram_id, requests, last_reset) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.TelegramID, user.Requests, user.LastReset)
	return err
}

func (r *UserRepository) ResetUserRequests(user *domain.User) {
	user.Requests = 0
	user.LastReset = time.Now()
	err := r.UpdateUserRequests(user)
	if err != nil {
		log.Println("Failed to reset user requests:", err)
	}
}
