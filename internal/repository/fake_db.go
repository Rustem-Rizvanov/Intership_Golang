package repository

import (
	"database/sql"
	"errors"
	"time"
	"Golang_Intership/internal/domain"
)

type FakeDB struct {
	users map[int64]*domain.User
}

func NewFakeDB() *FakeDB {
	return &FakeDB{
		users: make(map[int64]*domain.User),
	}
}

func (db *FakeDB) QueryRow(query string, args ...interface{}) *sql.Row {
	telegramID := args[0].(int64)
	user, exists := db.users[telegramID]
	if !exists {
		return &sql.Row{}
	}
	row := &sql.Row{}
	return row
}

func (db *FakeDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	switch query {
	case "INSERT INTO users":
		telegramID := args[0].(int64)
		db.users[telegramID] = &domain.User{
			ID:         1,
			TelegramID: telegramID,
			Requests:   args[1].(int),
			LastReset:  args[2].(time.Time),
		}
	case "UPDATE users":
		telegramID := args[0].(int64)
		user, exists := db.users[telegramID]
		if !exists {
			return nil, errors.New("user not found")
		}
		user.Requests = args[1].(int)
		user.LastReset = args[2].(time.Time)
	}
	return nil, nil
}

func (db *FakeDB) CreateUser(user *domain.User) error {
	db.users[user.TelegramID] = user
	return nil
}

func (db *FakeDB) GetUserByTelegramID(telegramID int64) (*domain.User, error) {
	user, exists := db.users[telegramID]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
