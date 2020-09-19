-include .env

VERSION:= $(shell git describe --tags)

install: 
	go get
start: 
	go run ./server.go
generate: 
	go run github.com/99designs/gqlgen generate
build: 
	go build -o bin/main .
compile:
	# Compiling for linux OS and platforms
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go