package domain

type HashRequest struct {
	Input string
}

type HashService interface {
	FindMD5Hash(input string) string
	CrackMD5Hash(hash string) string
}
