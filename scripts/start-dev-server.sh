#!/bin/bash
set -euo pipefail

echo -e "Start dev server.\n"

cleanup() {
    echo "Stopping services..."

    for pid in "${OATHKEEPER_PID:-}" "${KRATOS_PID:-}" "${HYDRA_PID:-}"; do
        if [ -n "${pid:-}" ] && kill -0 "$pid" 2>/dev/null; then
            kill "$pid"
            wait "$pid" 2>/dev/null || true
        fi
    done

    echo "Services stopped. Stop docker."

    docker compose -f ./dev/docker-compose.dev.yaml stop

    echo "Stopped."
}

trap cleanup EXIT INT TERM

# ========================
#       Before start
# ========================

echo "start dev docker compose"
docker compose -f ./dev/docker-compose.dev.yaml up -d

echo "apply migrations for kratos"
kratos migrate -c ./dev/config/kratos.yaml sql -e --yes

echo "apply migrations for hydra"
hydra migrate -c ./dev/config/hydra.yaml sql up -e --yes

echo "migrate application schema"
make migrate_up

# ========================
#       Start services
# ========================

echo "Starting Oathkeeper..."
oathkeeper serve proxy -c ./dev/config/oathkeeper.yaml &
OATHKEEPER_PID=$!

echo "Starting Kratos..."
kratos serve -c ./dev/config/kratos.yaml --dev --watch-courier &
KRATOS_PID=$!

echo "Starting Hydra..."
hydra serve -c ./dev/config/hydra.yaml all --dev &
HYDRA_PID=$!

wait
