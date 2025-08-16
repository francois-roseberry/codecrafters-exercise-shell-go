deps:
	go mod download

lint:
	gofmt -w -s .

test:
	go test -v -cover ./...

build:
	go build -o shell main.go

start:
	go run main.go

dev:
	air