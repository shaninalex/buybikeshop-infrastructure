import logging
import os

import databases
import sqlalchemy

DEBUG = os.getenv("DEBUG", default=False)

if DEBUG:
    # Enable SQL logging
    logging.basicConfig()
    logging.getLogger("databases").setLevel(logging.DEBUG)
    logging.getLogger("sqlalchemy.engine").setLevel(logging.INFO)

_connection_string: str = os.getenv(
    "AUTHSERVER_DB_CONNECTION_STRING",
    "postgresql://postgres:postgres@localhost:5432/postgres",
)

# Define metadata and database connection
metadata = sqlalchemy.MetaData()
database = databases.Database(_connection_string)
