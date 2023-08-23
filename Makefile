postgres:
	docker run --name menu-creator-db -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it menu-creator-db createdb --username=postgres --owner=postgres menu-creator-db

dropdb:
	docker exec -it menu-creator-db dropdb menu-creator-db

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/menu-creator-db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/menu-creator-db?sslmode=disable" -verbose down

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown server