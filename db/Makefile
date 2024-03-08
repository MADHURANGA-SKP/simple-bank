postgres:
	docker run -d --name bank -p 5432:5432 -e POSTGRES_USER=pasan -e POSTGRES_PASSWORD=12345 postgres:16-alpine

createdb:
	docker exec -it bank createdb --username=pasan --owner=pasan bank

dropdb:
	docker exec -it bank dropdb bank

migrateup:
	migrate -path db/migrations -database "postgresql://pasan:12345@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://pasan:12345@localhost:5432/bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc