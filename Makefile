postgres:
	docker run --name menu-creator-db --network menucreator-network-p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it menu-creator-postgres-1 createdb --username=postgres --owner=postgres menu-creator-db

dropdb:
	docker exec -it menu-creator-postgres-1 dropdb menu-creator-db

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/menu-creator-db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@postgres:5432/menu-creator-db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

# used to run the menucreator image that is built from Dockerfile
# can remove since using docker-compose file
menucreator:
	docker run --name menucreator --network menucreator-network -p 8080:8080 -e DB_SOURCE="postgresql://postgres:password@menu-creator-db:5432/menu-creator-db?sslmode=disabled" menucreator:latest

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server