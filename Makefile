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

.PHONY: test,createtstdb,mock

createtstdb:
	createdb apicl_test
	migrate -path migration -database "postgres://azunai:0000@localhost/apicl_test?sslmode=disable" up

test:
	go test -v -race -timeout 30s ./...

mock:
	mockgen -package mockdb -destination internal/repository/mock/mockdb.go goCleanArch/internal/repository Repository

.DEFAULT_GOAL := build