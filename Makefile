APP_NAME=go-microservice
BINARY_NAME=server

.PHONY: run build tidy docker-up docker-build docker-run clean

# Run locally
run:
	@echo "🚀 Running $(APP_NAME)..."
	go run cmd/server/main.go

# Build binary
build:
	@echo "🏗️  Building binary..."
	go build -o $(BINARY_NAME) ./cmd/server

# Format and tidy modules
tidy:
	go mod tidy
	go fmt ./...

# Build Docker image
docker-build:
	@echo "🐳 Building Docker image..."
	docker build -t $(APP_NAME):latest .

# Run Docker container
docker-run:
	@echo "🐳 Running Docker container..."
	docker run -d -p 8080:8080 --env-file .env $(APP_NAME):latest

# Stop and remove Docker containers
docker-stop:
	@echo "🧹 Stopping containers..."
	docker stop $$(docker ps -q --filter ancestor=$(APP_NAME):latest) || true

# Clean up
clean:
	@echo "🧼 Cleaning up..."
	rm -f $(BINARY_NAME)
