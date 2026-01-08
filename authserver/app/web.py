from starlette.applications import Starlette
from starlette.types import ASGIApp

from .config import Config
from .di import bootstrap

bootstrap()


def new_auth_server(config: Config) -> ASGIApp:
    return Starlette(
        debug=config.debug,  # TODO: get from settings
        routes=[],
        # lifespan=lifespan,
        # middleware=[],  # TODO: csrf, cookie
        # exception_handlers={
        #     404: not_found,
        # },
    )
