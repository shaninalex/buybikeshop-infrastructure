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
    drop_stmt = text("""DELETE FROM catalog.categories""")
    insert_stmt = text("""
        INSERT INTO catalog.categories (title, parent_id)
        VALUES (:title, :parent_id) RETURNING id
    """)

    def __init__(self, engine, system_path):
        self.system_path = system_path
        self.engine = engine

    def clear(self):
        with self.engine.connect() as conn:
            conn.execute(self.drop_stmt)
            conn.commit()

    def create(self):
        with open(f"{self.system_path}/../resources/categories.json") as f:
            d = json.load(f)
            print("Create categories")
            with self.engine.connect() as conn:
                self._make(conn, d, 0)
                conn.commit()

    def _make(self, conn, data: dict, parent_id: int | None = None):
        for d in data:
            result = conn.execute(
                self.insert_stmt,
                {
                    "title": d,
                    "parent_id": parent_id,
                },
            )
            cid = result.fetchone()[0]
            print(f"{cid}: {d}")
            if len(data[d].keys()) > 0:
                self._make(conn, data[d], cid)
