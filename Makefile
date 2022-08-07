APP_NAME=restapi-with-opentelemetry
IMAGE_TAG=$(shell git rev-parse --short HEAD)

build:
	go build -o deploy/bin/$(APP_NAME) cmd/server/main.go
test:
	go test ./... -cover -vet -all
docker-build:
	docker build -t $(APP_NAME):$(IMAGE_TAG) deploy/build
