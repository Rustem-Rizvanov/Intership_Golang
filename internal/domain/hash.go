package domain

type HashRequest struct {
	Input string
}

type HashService interface {
	HandleMessage(message string) (string, error) 
}
