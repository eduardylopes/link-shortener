.PHONY: run postgresinit postgres createdb dropdb migrateup migratedown

# Variables
APP_NAME=link-shortner
DATABASE_CONTAINER_NAME=link-shortener-db

# Tasks
default: run

run:
	@go run cmd/api/main.go

postgresinit:
	@docker run --name ${DATABASE_CONTAINER_NAME} -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

postgres:
	@docker exec -it ${DATABASE_CONTAINER_NAME} psql

createdb:
	@docker exec -it ${DATABASE_CONTAINER_NAME} createdb --username=root --owner=root ${APP_NAME}

dropdb:
	@docker exec -it ${DATABASE_CONTAINER_NAME} dropdb ${APP_NAME}

migrateup:
	@migrate -path configs/db/migrations -database "postgresql://root:password@localhost:5432/${APP_NAME}?sslmode=disable" -verbose up

migratedown:
	@migrate -path configs/db/migrations -database "postgresql://root:password@localhost:5432/${APP_NAME}?sslmode=disable" -verbose down
