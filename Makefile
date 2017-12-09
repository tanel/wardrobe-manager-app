default: build

build:
	go build

run: build
	./wardrobe-manager-app
