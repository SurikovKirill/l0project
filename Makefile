.PHONY:
build:
	go build -o ./.bin/bot cmd/main.go

run: build
	./.bin/l0pr
