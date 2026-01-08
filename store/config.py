import os
from dataclasses import dataclass


class ConfigVariableNotProvidedException(Exception):
    "Config variable was not provided"

    pass


@dataclass()
class Config:
    ClientSecret: str
    ClientId: str


def newConfig() -> Config:
    if os.getenv("STORE_CLIENT_SECRET", "") == "":
        raise ConfigVariableNotProvidedException(
            'Variable "STORE_CLIENT_SECRET" is empty'
        )

    if os.getenv("STORE_CLIENT_ID", "") == "":
        raise ConfigVariableNotProvidedException('Variable "STORE_CLIENT_ID" is empty')

    return Config(
        ClientSecret=os.getenv("STORE_CLIENT_SECRET", ""),
        ClientId=os.getenv("STORE_CLIENT_ID", ""),
    )
