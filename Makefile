include .env

MIGRATE_DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

migrate-up:
	migrate -path migrations -database "$(MIGRATE_DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(MIGRATE_DB_URL)" down 1

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

dev:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

