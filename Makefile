# Makefile for Go API project

# Commands
.PHONY: run
run:
	@echo "Running the application..."
	go run ./cmd/main.go

.PHONY: build
build:
	@echo "Building the application..."
	go build -o go-api.exe ./cmd/main.go

.PHONY: clean
clean:
	@echo "Cleaning up..."
	go clean
	rm .\go-api.exe

.PHONY: docker-build
docker-build:
	@echo "Building the Docker image..."
	docker compose up -d