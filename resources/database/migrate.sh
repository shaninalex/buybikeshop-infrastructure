#!/bin/bash

# Ensure required environment variables are set
if [[ -z "$DB_HOST" || -z "$DB_PORT" || -z "$DB_NAME" || -z "$DB_USER" || -z "$DB_PASSWORD" ]]; then
  echo "Error: Required database environment variables are not set."
  echo "Please set DB_HOST, DB_PORT, DB_NAME, DB_USER, and DB_PASSWORD."
  exit 1
fi


# migrate schema
/root/migrate \
		-path /root/migrations/ \
		-database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" \
		-verbose up


# fill db with defaults and required data
export PGPASSWORD="$DB_PASSWORD"
until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME"; do
  sleep 2
done

psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -f /root/default_data/defaults.sql

echo "defaults inserted."

# Add next script here with next pattern:
# psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -f /root/default_data/<script_name>.sql
