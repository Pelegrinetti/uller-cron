start:
	@echo "Running..."
	@go run cmd/main.go

build:
	@echo "Building..."
	@env CGO_ENABLED=1 GOARCH=arm GOOS=linux go build -o uller cmd/main.go