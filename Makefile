build:
	go build -o bin/$(shell basename $(PWD)) main.go

run:
	go run main.go

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(shell basename $(PWD)) main.go

window:
	GOOS=windows GOARCH=amd64 go build -o bin/$(shell basename $(PWD)).exe main.go

macos:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(shell basename $(PWD)) main.go

clean:
	rm -rf bin/*

all: linux window macos

.PHONY: build linux window macos

