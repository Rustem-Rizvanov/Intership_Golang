package domain

import "time"

type User struct {
	ID         int
	TelegramID int64
	Requests   int
	LastReset  time.Time
}
