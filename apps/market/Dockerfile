# Stage 1: Build dependencies
FROM python:3.14-slim AS build

COPY --from=ghcr.io/astral-sh/uv:latest /uv /usr/local/bin/uv

WORKDIR /app

# Copy workspace configuration for dependency resolution
COPY pyproject.toml uv.lock .python-version ./
COPY apps/market/pyproject.toml ./apps/market/pyproject.toml
COPY gen/grpc_buybikeshop_python/pyproject.toml ./gen/grpc_buybikeshop_python/pyproject.toml
COPY libs/python/common_tools/pyproject.toml ./libs/python/common_tools/pyproject.toml

# Install external dependencies first (cached layer)
RUN uv sync --frozen --no-dev --no-install-workspace --package market

# Copy source code for workspace packages
COPY apps/market ./apps/market
COPY gen/grpc_buybikeshop_python ./gen/grpc_buybikeshop_python
COPY libs/python/common_tools ./libs/python/common_tools

# Install workspace packages as non-editable
RUN uv sync --frozen --no-dev --no-editable --package market

# Stage 2: Runtime
FROM python:3.14-slim

WORKDIR /app

# Copy virtual environment with all packages installed
COPY --from=build /app/.venv /app/.venv

# Copy web assets (not included in Python package)
COPY apps/market/market/web/templates ./templates
COPY apps/market/market/web/static ./static

ENV PATH="/app/.venv/bin:$PATH"

EXPOSE 8004

CMD ["uvicorn", "market.main:app", "--host", "0.0.0.0", "--port", "8004"]
