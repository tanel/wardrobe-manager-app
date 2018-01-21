DATE=`date +%Y%m%d_%H%M%S`

default: build

build:
	go build

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -o wardrobe-linux

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

testuser: testuser-up

testuser-up:
	migrate -path testdata/migrations -url "postgres://wardrobe@localhost/wardrobe?sslmode=disable" up 1
	mkdir -p uploads/18f25d1b-dd0a-4889-9610-d103164c2f2e/item-images
	cp testdata/item-images/* uploads/18f25d1b-dd0a-4889-9610-d103164c2f2e/item-images

testuser-down:
	migrate -path testdata/migrations -url "postgres://wardrobe@localhost/wardrobe?sslmode=disable" down 1

backup:
	@mkdir -p backups
	@pg_dump wardrobe > backups/$(DATE).sql
	cp -r uploads backups/uploads

lint: lint-go lint-js lint-css

lint-go:
	gometalinter ./... --config=.gometalinter

lint-js:
	jshint public/js/app.js

lint-css:
	csslint public/css/app.css

thumbnails: 
	go run cmd/thumbnails/thumbnails.go

test:	
	TEMPLATE_PATH=../../template/*.html go test ./...

deploy: build-linux
	scp -r public deploy@wardrobe:/home/deploy/wardrobe/
	scp -r migrations deploy@wardrobe:/home/deploy/wardrobe/
	scp wardrobe-linux deploy@wardrobe:/home/deploy/wardrobe/
