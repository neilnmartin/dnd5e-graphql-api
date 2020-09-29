-include .env

# VERSION:= $(shell git describe --tags)

## TODO add dir watch

install: 
	go get
gen: 
	go run github.com/99designs/gqlgen generate
start: 
	go run ./server.go
run:
	./bin/server
compile:
	# Compiling for linux OS and platforms
	go build -o ./bin/server .
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 .
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm .
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 .
all: install gen compile run