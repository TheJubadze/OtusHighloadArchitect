postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-alpine

createdb:
	docker exec -it postgres16 createdb --username=root peepl

dropdb:
	docker exec -it postgres16 dropdb peepl

lint-fix:
	go mod tidy
	gofumpt -l -w .
	golangci-lint run --fix

docker-build:
	 docker compose up --build

.PHONY: postgres createdb dropdb lint-fix docker-build