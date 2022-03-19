.PHOHY: build, clear, up,stop,down

build:
	go build -v ./cmd/apiserver

clear:
	rm -f apiserver
up:
	docker-compose up -d
stop:
	docker-compose stop

down:
	docker-compose down


.PHONY: test
test:
	go test -v -race -timeout 30s ./...
.DEFAULT_GOAL := build