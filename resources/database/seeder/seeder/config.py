from dataclasses import dataclass
from pathlib import Path

import yaml


@dataclass
class Config:
    database: str


def read_config(path: Path) -> Config:
    with open(path, "r") as f:
        fconf = yaml.safe_load(f)
        return Config(**fconf)
