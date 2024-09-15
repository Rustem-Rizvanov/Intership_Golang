package repository

import (
	"Golang_Intership/internal/domain"
	"database/sql"
	"time"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByTelegramID(telegramID int64) (*domain.User, error) {
    user := &domain.User{}
    query := `SELECT id, telegram_id, requests, last_reset FROM users WHERE telegram_id = $1`
    err := r.db.QueryRow(query, telegramID).Scan(&user.ID, &user.TelegramID, &user.Requests, &user.LastReset)
    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) CreateUser(user *domain.User) error {
    query := `INSERT INTO users (telegram_id, requests, last_reset) VALUES ($1, $2, $3)`
    _, err := r.db.Exec(query, user.TelegramID, user.Requests, user.LastReset)
    return err
}

func (r *UserRepository) UpdateUserRequests(user *domain.User) error {
    query := `UPDATE users SET requests = $1, last_reset = $2 WHERE id = $3`
    _, err := r.db.Exec(query, user.Requests, user.LastReset, user.ID)
    return err
}

func (r *UserRepository) ResetUserRequests(user *domain.User) error {
    user.Requests = 0
    user.LastReset = time.Now()
    return r.UpdateUserRequests(user)
}
