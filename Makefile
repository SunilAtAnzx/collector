test:
	go test -v -cover -short ./...

server:
	go run main.go

docker-image:
	docker build -t go-collector-img .
.PHONY: test server docker-image
