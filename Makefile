default: build

build:
	go build

run: build
	./wardrobe-manager-app

migration:
	migrate -path migrations -url "postgres://wardrobe@localhost/wardrobe?sslmode=disable" create $(name)

migrate:
	migrate -path migrations -url "postgres://wardrobe@localhost/wardrobe?sslmode=disable" up

migrate-up:
	migrate -path migrations -url "postgres://wardrobe@localhost/wardrobe?sslmode=disable" up 1

migrate-down:
	migrate -path migrations -url "postgres://wardrobe@localhost/wardrobe?sslmode=disable" down 1
