#!/bin/bash

echo -e "Start warehouse API server.\n"

go run apps/warehouse/app/main.go serve --config=./apps/warehouse/config/config.dev.yaml
