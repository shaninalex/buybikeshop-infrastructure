# usage:
# 	make migrate_up name=test - create migration called "<time in "20060102150405" format>_test.(up/down).sql
migrate_create:
	~/go/bin/migrate create -ext sql -dir ./database/migrations -format "20060102150405" $(name)

migrate_up:
	~/go/bin/migrate \
		-path ./database/migrations/ \
		-database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" \
		-verbose up

# usage:
# 	make migrate_down N=1 - for one migration down
migrate_down:
	~/go/bin/migrate \
		-path ./database/migrations/ \
		-database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" \
		-verbose down $(N)

start:
	docker compose \
	    -f ./docker/docker-compose.base.yaml \
		-f ./docker/hydra.docker.yaml \
		-f ./docker/kratos.docker.yaml \
		up -d --build

stop:
	docker compose \
        -f ./docker/docker-compose.base.yaml \
        -f ./docker/hydra.docker.yaml \
        -f ./docker/kratos.docker.yaml \
		stop

clear:
	docker compose \
        -f ./docker/docker-compose.base.yaml \
        -f ./docker/hydra.docker.yaml \
        -f ./docker/kratos.docker.yaml \
		down -v
