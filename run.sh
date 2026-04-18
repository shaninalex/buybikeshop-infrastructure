#!/usr/bin/env bash

set -e

PROJECT_ROOT="$(cd "$(dirname "$0")" && pwd)"

COMPOSE_MAIN=(
  -f "$PROJECT_ROOT/docker/docker-compose.base.yml"
  -f "$PROJECT_ROOT/docker/datasource.docker.yml"
#  -f "$PROJECT_ROOT/docker/warehouse.docker.yml"
#  -f "$PROJECT_ROOT/docker/market.docker.yml"
  -f "$PROJECT_ROOT/docker/kratos.docker.yml"
  -f "$PROJECT_ROOT/docker/keto.docker.yml"
#  -f "$PROJECT_ROOT/docker/office.docker.yml"
#  -f "$PROJECT_ROOT/docker/admin.docker.yml"
)

COMPOSE_TEST=(
  -f "$PROJECT_ROOT/tdata/docker-compose.test.yml"
)

MIGRATE_BIN="${MIGRATE_BIN:-$HOME/go/bin/migrate}"

DB_URL="${DB_URL:-postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable}"

function migrate_create() {
  local name="$1"

  if [[ -z "$name" ]]; then
    echo "Usage: ./run.sh migrate_create <name>"
    exit 1
  fi

  "$MIGRATE_BIN" create \
    -ext sql \
    -dir "$PROJECT_ROOT/database/migrations" \
    -format "20060102150405" "$name"
}

function migrate_up() {
  "$MIGRATE_BIN" \
    -path "$PROJECT_ROOT/database/migrations" \
    -database "$DB_URL" \
    -verbose up
}

function migrate_down() {
  local steps="${1:-1}"

  "$MIGRATE_BIN" \
    -path "$PROJECT_ROOT/database/migrations" \
    -database "$DB_URL" \
    -verbose down "$steps"
}

function start() {
  docker compose "${COMPOSE_MAIN[@]}" up -d --build
}

function stop() {
  docker compose "${COMPOSE_MAIN[@]}" stop
}

function clear() {
  docker compose "${COMPOSE_MAIN[@]}" down -v
}

function rebuild() {
  local service="$1"

  if [[ -z "$service" ]]; then
    echo "Usage: ./run.sh rebuild <service>"
    exit 1
  fi

  docker compose "${COMPOSE_MAIN[@]}" up -d --no-deps --build "$service"
}

function start_db() {
  docker compose -f "$PROJECT_ROOT/docker/docker-compose.base.yml" up -d --build
}

function clear_db() {
  docker compose -f "$PROJECT_ROOT/docker/docker-compose.base.yml" down -v
}

function generate_go_grpc() {
  mkdir -p "$PROJECT_ROOT/gen"

  protoc -I "$PROJECT_ROOT/proto" \
    --go_out="$PROJECT_ROOT/gen" \
    --go-grpc_out="$PROJECT_ROOT/gen" \
    $(find "$PROJECT_ROOT/proto" -name '*.proto')
}

function generate_python_grpc() {
  mkdir -p "$PROJECT_ROOT/gen/grpc_buybikeshop_python"

  uv run python -m grpc_tools.protoc \
    -I"$PROJECT_ROOT/proto" \
    --python_out="$PROJECT_ROOT/gen/grpc_buybikeshop_python" \
    --pyi_out="$PROJECT_ROOT/gen/grpc_buybikeshop_python" \
    --grpc_python_out="$PROJECT_ROOT/gen/grpc_buybikeshop_python" \
    $(find "$PROJECT_ROOT/proto" -name '*.proto')

  find "$PROJECT_ROOT/gen/grpc_buybikeshop_python" -type d -exec touch {}/__init__.py \;
}

function generate_grpc() {
  generate_go_grpc
  generate_python_grpc
  echo "Completed."
}

function start_test_db() {
  docker compose "${COMPOSE_TEST[@]}" up -d --build
}

function clear_test_db() {
  docker compose "${COMPOSE_TEST[@]}" down -v
}

function help() {
  echo "Available commands:"
  echo "  migrate_create <name>"
  echo "  migrate_up"
  echo "  migrate_down [N]"
  echo "  start | stop | clear | rebuild <service>"
  echo "  start_db | clear_db"
  echo "  start_test_db"
  echo "  generate_go_grpc"
  echo "  generate_python_grpc"
  echo "  generate_grpc"
}

function seed() {
   uv run --package seeder python -m seeder.main \
    --config ./database/seeder/config.yaml \
    start
}

# --- Dispatcher ---
cmd="$1"
shift || true

case "$cmd" in
  migrate_create) migrate_create "$@" ;;
  migrate_up) migrate_up ;;
  migrate_down) migrate_down "$@" ;;
  start) start ;;
  stop) stop ;;
  clear) clear ;;
  rebuild) rebuild "$@" ;;
  start_db) start_db ;;
  clear_db) clear_db ;;
  start_test_db) start_test_db ;;
  clear_test_db) clear_test_db ;;
  generate_go_grpc) generate_go_grpc ;;
  generate_python_grpc) generate_python_grpc ;;
  generate_grpc) generate_grpc ;;
  seed) seed ;;
  ""|help|-h|--help) help ;;
  *)
    echo "Unknown command: $cmd"
    help
    exit 1
    ;;
esac
