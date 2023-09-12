postgres:
	docker run --name postgres12-d -p 5436:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12-d createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12-d dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5436/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5436/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
	