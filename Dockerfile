# Этап сборки
FROM golang:1.23.0-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o telegrambot cmd/telegrambot/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/telegrambot .
COPY .env .env
CMD ["./telegrambot"]
