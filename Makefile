BINARY_NAME=collector
LINUX_BINARY_NAME=linux_collector

test:
	go test -v -cover -short ./...

server:
	go run main.go -p 8282

docker-image:
	docker build -t go-collector-img .

build: bin/$(BINARY_NAME)

bin/$(BINARY_NAME):
	go build -o bin/$(BINARY_NAME) .
	env GOOS=linux GOARCH=amd64 go build -o bin/$(LINUX_BINARY_NAME) .

.PHONY: build test server docker-image
