package domain

type UserRepository interface {
    GetUserByTelegramID(telegramID int64) (*User, error)
    CreateUser(user *User) error
    UpdateUserRequests(user *User) error
    ResetUserRequests(user *User) error
}
