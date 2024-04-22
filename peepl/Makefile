postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-alpine

createdb:
	docker exec -it postgres16 createdb --username=root peepl

dropdb:
	docker exec -it postgres16 dropdb peepl

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/peepl?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/peepl?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown