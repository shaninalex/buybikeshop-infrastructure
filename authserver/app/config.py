from dataclasses import dataclass

import yaml


@dataclass
class Config:
    database_url: str
    debug: bool
    port: int


def load_config(path: str) -> Config:
    with open(path, "r") as f:
        data = yaml.safe_load(f)

    return Config(
        database_url=data["database"]["url"],
        debug=data.get("debug", False),
        port=data.get("port", 8000),
    )
