.PHONY: build install

build:
	go build -o bin/cog -trimpath -ldflags="-s -w" main.go

install: build
	sudo install bin/cog /usr/local/bin
