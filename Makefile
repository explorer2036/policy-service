default: build

all: build

build:
	go build -o policy-server main.go
