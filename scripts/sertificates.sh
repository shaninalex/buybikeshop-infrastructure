#!/bin/bash

# client keys
openssl genrsa -out ca-key.pem 4096
openssl req -x509 -new -nodes \
  -key ca-key.pem \
  -sha256 \
  -days 3650 \
  -out ca.pem \
  -subj "/CN=buybikeshop-ca"

# server keys
openssl genrsa -out server-key.pem 4096

openssl req -new \
  -key server-key.pem \
  -out server.csr \
  -subj "/CN=localhost"

openssl x509 -req \
  -in server.csr \
  -CA ca.pem \
  -CAkey ca-key.pem \
  -CAcreateserial \
  -out server.pem \
  -days 365 \
  -sha256
