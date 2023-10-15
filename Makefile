DB_URL=postgresql://postgres:password@localhost:5435/menu-creator-db?sslmode=disable

postgres:
	docker run --name menu-creator-db --network menucreator-network-p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it menu-creator-postgres-1 createdb --username=postgres --owner=postgres menu-creator-db

dropdb:
	docker exec -it menu-creator-postgres-1 dropdb menu-creator-db

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server test 