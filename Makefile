run:
	go run ./cmd/server

test:
	go test ./...

build:
	go build -o fizzbuzz ./cmd/server

lint:
	golangci-lint run
