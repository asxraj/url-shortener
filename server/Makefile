include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# Create the new confirm target.
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
run/api:
	go run ./cmd/api

## test: run the cmd/api tests with -v flag
test:
	go test -v ./...

## psql: connect to the database using psql
psql:
	docker exec -it postgres14 psql ${SHORTURL_DB_DSN}
## postgres: run image container 
postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pa55word -d postgres:14-alpine

## run redis image container
redis:
	docker run -d --name redis-db -p 6379:6379 redis

## createdb: create shorturl database
createdb:
	docker exec -it postgres14 createdb shorturl

## startdb: run postgres container and create db
startdb: 
	make postgres
	make createdb

drop:
	docker stop postgres14
	docker rm postgres14

## dropdb: drop database shorturl
dropdb:
	docker exec -it postgres14 dropdb shorturl

## db/migrations/up: apply all up database migrations
migrate/up: 
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${SHORTURL_DB_DSN} up

## db/migrations/down: apply all down database migrations
migrate/down:
	@echo 'Running down migrations...'
	migrate -path ./migrations -database ${SHORTURL_DB_DSN} down

again:
	docker stop postgres14
	docker rm postgres14
	make postgres

## rebuild: deleting docker image and rebuilding	
rebuild: 
	@echo 'Docker compose down'
	docker compose down
# @echo 'Removing docker image'
# docker rmi server-api
	@echo 'Docker compose up'
	docker compose up -d

 
.PHONY: help confirm run/api test psql postgres createdb dropdb migrate/up migrate/down 