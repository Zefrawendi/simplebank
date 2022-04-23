postgres:
	docker run -itd --name postgres -p 5432:5432 \
-e POSTGRES_USER=devops -e POSTGRES_PASSWORD="devops" postgres

createdb:
	docker exec -it postgres createdb --username devops --owner=devops simple_bank

dropdb:
	docker exec -it postgres dropdb --username devops simple_bank

migrateup:
	migrate -path db/migration --database "postgresql://devops:devops@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration --database "postgresql://devops:devops@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test