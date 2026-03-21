# usage:
# 	make migrate_up name=test - create migration called "<time in "20060102150405" format>_test.(up/down).sql
migrate_create:
	~/go/bin/migrate create \
		-ext sql \
		-dir ./database/migrations \
		-format "20060102150405" $(name)

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
	    -f ./docker/docker-compose.base.yml \
		-f ./docker/datasource.docker.yml \
		-f ./docker/warehouse.docker.yml \
		-f ./docker/market.docker.yml \
		up -d --build

stop:
	docker compose \
        -f ./docker/docker-compose.base.yml \
		-f ./docker/datasource.docker.yml \
		-f ./docker/warehouse.docker.yml \
		-f ./docker/market.docker.yml \
		stop

clear:
	docker compose \
        -f ./docker/docker-compose.base.yml \
		-f ./docker/datasource.docker.yml \
		-f ./docker/warehouse.docker.yml \
		-f ./docker/market.docker.yml \
		down -v


rebuild:
	docker compose \
        -f ./docker/docker-compose.base.yml \
		-f ./docker/datasource.docker.yml \
		-f ./docker/warehouse.docker.yml \
		-f ./docker/market.docker.yml \
		up -d --no-deps --build $(name)

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

generate_python_grpc:
	uv run python -m grpc_tools.protoc -I./proto \
		--python_out=./gen/grpc_buybikeshop_python \
		--pyi_out=./gen/grpc_buybikeshop_python \
		--grpc_python_out=./gen/grpc_buybikeshop_python \
		$$(find proto -name '*.proto')
	find ./gen/grpc_buybikeshop_python -type d -exec touch {}/__init__.py \;

generate_grpc: generate_go_grpc generate_python_grpc