default: build

all: build

build:
	go build -o policy-service main.go
