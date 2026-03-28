import os

import grpc

from common_tools.di import Container
from market.core.connector import CatalogConnector
from market.web.templating import Templates


def _configure(container: Container) -> None:
    channel = grpc.insecure_channel(os.environ.get("APP_MARKET_DATASOURCE"))
    container.register(CatalogConnector, instance=CatalogConnector(channel))
    container.register(Templates, instance=Templates())


def create_container() -> Container:
    return Container(_configure)


_container = create_container()
bootstrap = _container.bootstrap
resolve = _container.resolve
