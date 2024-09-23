package adapter

import (
    "github.com/streadway/amqp"
    "log"
    "time"
	"fmt"
)

type RabbitMQAdapter struct {
    connection *amqp.Connection
}

func NewRabbitMQAdapter(url string) (*RabbitMQAdapter, error) {
    var conn *amqp.Connection
    var err error

    for {
        conn, err = amqp.Dial(url)
        if err == nil {
            log.Println("Connected to RabbitMQ successfully.")
            break
        }
        log.Printf("Failed to connect to RabbitMQ, retrying in 5 seconds: %s", err)
        time.Sleep(5 * time.Second) 
    }

    return &RabbitMQAdapter{connection: conn}, nil
}

func (r *RabbitMQAdapter) Close() error {
    if r.connection != nil {
        return r.connection.Close()
    }
    return nil
}

func (r *RabbitMQAdapter) Publish(queue string, message []byte) error {
    channel, err := r.connection.Channel()
    if err != nil {
        log.Printf("Failed to create channel: %v", err)
        return err
    }
    defer channel.Close()

    _, err = channel.QueueDeclare(queue, true, false, false, false, nil)
    if err != nil {
        log.Printf("Failed to declare queue: %v", err)
        return err
    }

    log.Printf("Publishing message to queue %s: %s", queue, message)
    err = channel.Publish(
        "", "telegram_queue", false, false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        message,
        },
    )
    if err != nil {
        log.Printf("Failed to publish message: %v", err)
    }
    return err
}

func (r *RabbitMQAdapter) SendTelegramMessage(telegramID int64, message string) error {
    msg := fmt.Sprintf(`{"telegram_id": %d, "text": "%s"}`, telegramID, message)
    return r.Publish("telegram_queue", []byte(msg))
}

func (r *RabbitMQAdapter) Consume(queue string, handler func([]byte) error) error {
    channel, err := r.connection.Channel()
    if err != nil {
        return fmt.Errorf("failed to open a channel: %w", err)
    }
    defer channel.Close()

    _, err = channel.QueueDeclare(queue, true, false, false, false, nil)
    if err != nil {
        return fmt.Errorf("failed to declare queue: %w", err)
    }

    msgs, err := channel.Consume(
        queue, 
        "",    
        true,  
        false, 
        false, 
        false, 
        nil,   
    )
    if err != nil {
        return fmt.Errorf("failed to start consuming: %w", err)
    }

    log.Printf("Consuming messages from queue %s", queue)
    
    for msg := range msgs {
        log.Printf("Received message: %s", msg.Body)
        if err := handler(msg.Body); err != nil {
            log.Printf("Error handling message: %v", err)
        } else {
            log.Printf("Successfully processed message: %s", msg.Body)
        }
    }

    log.Printf("No more messages in queue %s", queue)
    return nil
}
