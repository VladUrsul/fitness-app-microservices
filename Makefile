.PHONY: tidy swagger build up

tidy:
	cd user-service && go mod tidy
	cd workout-service && go mod tidy
	cd session-service && go mod tidy
	cd api-gateway && go mod tidy

swagger:
	cd api-gateway && swag init -g ./cmd/gateway/main.go -o ./internal/docs

build:
	docker-compose build

up:
	docker-compose up
