.PHONY: build
build:
	go build -v ./cmd/vpnbot
	./vpnbot

.PHONY: test
test:
	go test -v -race -timeout 30s ./...


.DEFAULT_GOAL := build