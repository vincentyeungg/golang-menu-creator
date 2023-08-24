postgres:
	docker run --name menu-creator-db --network menucreator-network-p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:12-alpine

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

# used to run the menucreator image that is built from Dockerfile
menucreator:
	docker run --name menucreator --network menucreator-network -p 8080:8080 -e DB_SOURCE="postgresql://postgres:password@menu-creator-db:5432/menu-creator-db?sslmode=disabled" menucreator:latest

.PHONY: postgres createdb dropdb migrateup migratedown server