start:
	@echo "Running..."
	@go run cmd/main.go

build:
	@echo "Building..."
	@go build -o uller cmd/main.go