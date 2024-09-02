build:
	go build -o bin/telegrambot cmd/telegrambot/main.go

run:
	go run cmd/telegrambot/main.go

test:
	go test ./... 


docker-build:
	docker build -t myproject .

docker-run:
	docker run --rm -e TELEGRAM_TOKEN=$(TELEGRAM_TOKEN) myproject
