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