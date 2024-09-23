package domain

type MessageBroker interface {
    Publish(queue string, message []byte) error
    Consume(queue string, handler func([]byte) error) error
}
