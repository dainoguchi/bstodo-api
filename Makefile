include .env.local

.PHONY: build
build:
	docker compose build --no-cache

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose dowm

#################
# マイグレーション #
#################

.PHONY: mgcreate
mgcreate:
	migrate create -ext sql -dir db/migrations -seq $(MG_NAME)

.PHONY: mgup
mgup:
	migrate -path db/migrations -database "$(DB_URL)" up

