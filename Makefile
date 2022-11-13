postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=toor -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root user_balances

dropdb:
	docker exec -it postgres15 dropdb user_balances

migrateup:
	migrate -path db/migration -database "postgresql://root:toor@localhost:5432/user_balances?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:toor@localhost:5432/user_balances?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server