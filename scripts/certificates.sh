#!/bin/bash

# TODO:
# Ideally each server should have own credentials files

set -e

# ---- configuration ----
CA_CN="${1:-buybikeshop-ca}"
SERVER_CN="${2:-datasource}"

CREDS_DIR="../creds"

# shellcheck disable=SC2164
cd ./scripts

# ---- client CA ----
openssl genrsa -out "$CREDS_DIR/ca-key.pem" 4096

openssl req -x509 -new -nodes \
    -key "$CREDS_DIR/ca-key.pem" \
    -sha256 \
    -days 3650 \
    -out "$CREDS_DIR/ca.pem" \
    -subj "/CN=${CA_CN}"

# ---- server keys ----
openssl genrsa -out "$CREDS_DIR/server-key.pem" 4096

openssl req -new \
    -key "$CREDS_DIR/server-key.pem" \
    -out "$CREDS_DIR/server.csr" \
    -subj "/CN=${SERVER_CN}"

openssl x509 -req \
    -in "$CREDS_DIR/server.csr" \
    -CA "$CREDS_DIR/ca.pem" \
    -CAkey "$CREDS_DIR/ca-key.pem" \
    -CAcreateserial \
    -out "$CREDS_DIR/server.pem" \
    -days 365 \
    -sha256
