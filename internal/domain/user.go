package domain

import "time"

type User struct {
    ID         int64     `json:"id"`
    TelegramID int64     `json:"telegram_id"`
    ChatID     int64     `json:"chat_id"`  
    Requests   int       `json:"requests"`
    LastReset  time.Time `json:"last_reset"`
}
