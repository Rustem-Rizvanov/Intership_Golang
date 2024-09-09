package service

import (
    "database/sql"
    "time"
    _ "github.com/lib/pq"
)

type UserService struct {
    db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
    return &UserService{db: db}
}

func (s *UserService) CheckAndUpdateUserRequests(telegramID int64, currentTime time.Time) (bool, error) {
    var requestsCount int
    var lastReset time.Time
    var banEnd sql.NullTime

    // Получаем данные о пользователе
    err := s.db.QueryRow(`
        SELECT requests_count, last_reset, ban_end
        FROM users
        WHERE telegram_id = $1`, telegramID).Scan(&requestsCount, &lastReset, &banEnd)
    if err != nil && err != sql.ErrNoRows {
        return false, err
    }

    // Если пользователь не найден, создаем его
    if err == sql.ErrNoRows {
        _, err := s.db.Exec(`
            INSERT INTO users (telegram_id, requests_count, last_reset, ban_end)
            VALUES ($1, $2, $3, $4)`, telegramID, 1, currentTime, nil)
        if err != nil {
            return false, err
        }
        return true, nil
    }

    // Проверяем время последнего сброса
    if currentTime.Sub(lastReset) > time.Hour {
        // Сброс количества запросов
        _, err := s.db.Exec(`
            UPDATE users
            SET requests_count = 1, last_reset = $1, ban_end = NULL
            WHERE telegram_id = $2`, currentTime, telegramID)
        if err != nil {
            return false, err
        }
        return true, nil
    }

    // Проверяем, забанен ли пользователь
    if banEnd.Valid && currentTime.Before(banEnd.Time) {
        return false, nil
    }

    if requestsCount >= 10 {
        _, err := s.db.Exec(`
            UPDATE users
            SET ban_end = $1
            WHERE telegram_id = $2`, currentTime.Add(10*time.Second), telegramID)
        if err != nil {
            return false, err
        }
        return false, nil
    }

    // Обновляем количество запросов
    _, err = s.db.Exec(`
        UPDATE users
        SET requests_count = requests_count + 1
        WHERE telegram_id = $1`, telegramID)
    if err != nil {
        return false, err
    }

    return true, nil
}
