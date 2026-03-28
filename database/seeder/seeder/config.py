from dataclasses import dataclass
from pathlib import Path

import yaml


@dataclass(frozen=True)
class PartnerConfig:
    amount: int

@dataclass(frozen=True)
class DatabaseConfig:
    database: str
    user: str
    password: str
    host: str
    port: int


@dataclass(frozen=True)
class Config:
    database: DatabaseConfig
    partners: PartnerConfig


def read_config(path: Path) -> Config:
    with open(path, "r") as f:
        fconf = yaml.unsafe_load(f)
        return Config(
            database=DatabaseConfig(**fconf['database']),
            partners=PartnerConfig(**fconf['partners']),
        )
