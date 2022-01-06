default: build

all: build

build:
	go build -o policy-service main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/alpine/policy-service main.go

docker:
	docker build -t policy-service:v1.0 .
