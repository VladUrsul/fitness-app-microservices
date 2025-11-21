.PHONY: tidy swagger build up

tidy:
	cd user-service && go mod tidy
	cd workout-service && go mod tidy
	cd session-service && go mod tidy

swagger:
	cd user-service && swag init -g ./cmd/server/main.go -o ./internal/docs
	cd workout-service && swag init -g ./cmd/server/main.go -o ./internal/docs
	cd session-service && swag init -g ./cmd/server/main.go -o ./internal/docs

build:
	docker-compose build

up:
	docker-compose up
