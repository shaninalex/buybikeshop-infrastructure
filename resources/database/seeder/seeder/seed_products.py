import json
from pathlib import Path

from sqlalchemy import create_engine, text

from seeder.config import read_config

path = Path(__file__).parent.resolve()


def seed_products(config: Path):
    conf = read_config(config)
    engine = create_engine(conf.database)
    c = CategoriesSeeder(engine, path)
    c.create()


class CategoriesSeeder:
    drop_stmt = text("""DELETE FROM warehouse.categories""")
    insert_stmt = text("""
        INSERT INTO warehouse.categories (title, level, parent_id)
        VALUES (:title, :level, :parent_id) RETURNING id
    """)

    def __init__(self, engine, path):
        self.path = path
        self.engine = engine

    def clear(self):
        with self.engine.connect() as conn:
            conn.execute(self.drop_stmt)
            conn.commit()

    def create(self):
        with open(f"{self.path}/../resources/categories.json") as f:
            d = json.load(f)
            print("Create categories")
            with self.engine.connect() as conn:
                self._make(conn, d, 0)
                conn.commit()

    def _make(self, conn, data: dict, level: int, parent_id: int | None = None):
        for d in data:
            result = conn.execute(
                self.insert_stmt,
                {
                    "title": d,
                    "level": level,
                    "parent_id": parent_id,
                },
            )
            cid = result.fetchone()[0]
            print(f"{cid}: {d}")
            if len(data[d].keys()) > 0:
                self._make(conn, data[d], level + 1, cid)
