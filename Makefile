APP_NAME=go-microservice

run:
	go run cmd/server/main.go

build:
	go build -o server ./cmd/server

docker-up:
	docker-compose up --build -d

docker-down:
	docker-compose down

tidy:
	go mod tidy
	go fmt ./...

clean:
	rm -f server
