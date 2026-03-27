import random
from typing import List

import psycopg2
from psycopg2 import OperationalError

from seeder.config import Config

DEFAULT_MEDIA = [
    "/img/bike0.jpg",
    "/img/bike1.jpg",
    "/img/bike2.jpg",
    "/img/bike3.jpg",
    "/img/bike4.jpg"
]


def get_random_media() -> List[str]:
    count = random.randint(0, len(DEFAULT_MEDIA))
    return random.sample(DEFAULT_MEDIA, count)


def create_connection(config: Config):
    connection = None
    try:
        connection = psycopg2.connect(
            database=config.database.database,
            user=config.database.user,
            password=config.database.password,
            host=config.database.host,
            port=config.database.port,
        )
        print("Connection successful")
    except OperationalError as e:
        raise OperationalError(f"The error '{e}' occurred")
    return connection


def execute_raw_sql(connection, sql):
    connection.autocommit = True
    cursor = connection.cursor()
    cursor.execute(sql)


def generate_sequence(start, end) -> List[int]:
    _end = random.randint(start, end)
    return list(range(1, _end + 1))