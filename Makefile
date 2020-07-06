start:
	@echo "Running..."
	@go run cmd/main.go

build:
	@echo "Building..."
	@env CGO_ENABLED=1 GOARCH=arm64 GOOS=linux go build -o uller cmd/main.go