# usage:
# 	make migrate_up name=test - create migration called "<time in "20060102150405" format>_test.(up/down).sql
migrate_create:
	~/go/bin/migrate create \
		-ext sql \
		-dir ./resources/database/migrations \
		-format "20060102150405" $(name)

migrate_up:
	~/go/bin/migrate \
		-path ./resources/database/migrations/ \
		-database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" \
		-verbose up

# usage:
# 	make migrate_down N=1 - for one migration down
migrate_down:
	~/go/bin/migrate \
		-path ./resources/database/migrations/ \
		-database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" \
		-verbose down $(N)

start:
	docker compose \
	    -f ./docker/docker-compose.base.yml \
		-f ./docker/hydra.docker.yml \
		-f ./docker/kratos.docker.yml \
		-f ./docker/oathkeeper.yml \
		-f ./docker/datasource.docker.yml \
		-f ./docker/warehouse.docker.yml \
		up -d --build

stop:
	docker compose \
        -f ./docker/docker-compose.base.yml \
        -f ./docker/hydra.docker.yml \
        -f ./docker/kratos.docker.yml \
		-f ./docker/oathkeeper.yml \
		-f ./docker/datasource.docker.yml \
		-f ./docker/warehouse.docker.yml \
		stop

clear:
	docker compose \
        -f ./docker/docker-compose.base.yml \
        -f ./docker/hydra.docker.yml \
        -f ./docker/kratos.docker.yml \
		-f ./docker/oathkeeper.yml \
		-f ./docker/datasource.docker.yml \
		-f ./docker/warehouse.docker.yml \
		down -v

start_db:
	docker compose \
        -f ./docker/docker-compose.base.yml \
        up -d --build

clear_db:
	docker compose \
        -f ./docker/docker-compose.base.yml \
		down -v

generate_go_grpc:
	protoc -I proto \
		--go_out=./gen \
		--go-grpc_out=./gen \
		$$(find proto -name '*.proto')
