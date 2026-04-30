from typing import Protocol, Dict, List

from seeder.config import Config
from seeder.entities import PartnersSeeder, DepartmentSeeder
from seeder.utils import create_connection


class Seeder(Protocol):
    def name(self) -> str: ...

    def seed(self, connection): ...

    def clear(self, connection): ...


class SeederRegistry:
    def __init__(self):
        self._registry: Dict[str, Seeder] = {}

    def register(self, s: Seeder):
        self._registry[s.name] = s

    def list(self) -> List[Seeder]:
        return self._registry.values()

def log(s: Seeder, msg: str):
    print(f"[{s.name()}]: {msg}")


class Executor:
    def __init__(self, config: Config, registry: SeederRegistry):
        self._config = config
        self._registry = registry

    def clear(self, connection):
        for s in self._registry.list():
            log(s, "clearing")
            s.clear(connection)

    def run(self, connection):
        for s in self._registry.list():
            log(s, "seeding...")
            s.seed(connection)
            log(s, "complete.")


def start(config: Config):
    db = create_connection(config)
    registry = SeederRegistry()
    registry.register(PartnersSeeder())
    registry.register(DepartmentSeeder())

    executor = Executor(config, registry)
    executor.clear(db)
    executor.run(db)
