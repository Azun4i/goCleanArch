.PHOHY: build, clear, up

build:
	go build -v ./cmd/apiserver

clear:
	rm -f apiserver
up:
	docker-compose up -d
stop:
	docker-compose stop

drop:
	docker-compose drop

.DEFAULT_GOAL := build