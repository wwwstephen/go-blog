postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root go_blog 

dropdb:
	docker exec -it postgres dropdb go_blog

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go_blog?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go_blog?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

run:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
