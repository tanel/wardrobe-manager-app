default: build

build:
	go build

run: build
	./wardrobe-manager-app

migration:
	migrate -path migrations -url "postgres://tanel@localhost/wardrobe?sslmode=disable" create $(name)

migrate:
	migrate -path migrations -url "postgres://tanel@localhost/wardrobe?sslmode=disable" up

migrate-down:
	migrate -path migrations -url "postgres://tanel@localhost/wardrobe?sslmode=disable" down
