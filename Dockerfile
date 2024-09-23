# Этап сборки
FROM golang:1.23.0-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o telegrambot cmd/telegrambot/main.go

# Этап сборки RabbitMQ
FROM rabbitmq:3-management as rabbitmq
# Этап запуска
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/telegrambot .
COPY .env .env

# Открываем порты (если нужны)
EXPOSE 8080

CMD ["./telegrambot"]
