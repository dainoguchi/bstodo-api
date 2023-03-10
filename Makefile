include .env

.PHONY: init
init:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	make build
	make up

.PHONY: build
build:
	docker compose build --no-cache

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose dowm

.PHONY: psql
psql:
	PGPASSWORD=$(DB_PASS) psql -h 127.0.0.1 -U $(DB_USER) -d $(DB_NAME)

.PHONY: sqlc-gen
sqlc-gen:
	docker compose run app sqlc generate

#################
# マイグレーション #
#################

.PHONY: mgcreate
mgcreate:
	migrate create -ext sql -dir db/migrations -seq $(MG_NAME)

.PHONY: mgup
mgup:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASS)@127.0.0.1:5432/$(DB_NAME)?sslmode=disable" up

.PHONY: mgup1
mgup1:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASS)@127.0.0.1:5432/$(DB_NAME)?sslmode=disable" up 1


.PHONY: mgdown
mgdown:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASS)@127.0.0.1:5432/$(DB_NAME)?sslmode=disable" down

.PHONY: mgdown1
mgdown1:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASS)@127.0.0.1:5432/$(DB_NAME)?sslmode=disable" down 1

.PHONY: mgforce
mgforce: ## error後必須
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASS)@127.0.0.1:5432/$(DB_NAME)?sslmode=disable" force $(MG_VERSION)

.PHONE: apply_test
apply_test:
	make mgup_test
	make seed_test

.PHONY: createdb_test
createdb_test:
	PGPASSWORD=$(DB_PASS) createdb -h 127.0.0.1 -U $(DB_USER) $(DB_NAME_TEST)

.PHONY: mgup_test
mgup_test:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASS)@127.0.0.1:5432/$(DB_NAME_TEST)?sslmode=disable" up

.PHONY: seed_test
seed_test:
	PGPASSWORD=$(DB_PASS) psql -h 127.0.0.1 -U $(DB_USER) -d $(DB_NAME_TEST) -f db/seeds/scenario.sql
