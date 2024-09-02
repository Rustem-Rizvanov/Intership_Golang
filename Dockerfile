FROM golang:1.23.0-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o telegrambot cmd/telegrambot/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/telegrambot .
CMD ["./telegrambot"]
