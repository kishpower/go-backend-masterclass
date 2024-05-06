postgres:
	docker-compose up -d

createdb:
	docker exec -it my-postgres-container createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it my-postgres-container dropdb simple_bank	

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5555/simple_bank?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5555/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
 	# uses reflection to create mocks
	mockgen -package mockdb -destination db/mock/store.go github.com/kishpower/simplebank/db/sqlc Store


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server