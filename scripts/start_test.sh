#!/bin/bash

docker compose \
    -f ./tdata/docker-compose.test.yml \
    up -d --build
