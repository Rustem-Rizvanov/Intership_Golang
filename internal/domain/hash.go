package domain

type HashRequest struct {
	Input string
}

type HashService interface {
	FindMD5Hash(input string) string
}
