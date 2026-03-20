# import grpc
from starlette.applications import Starlette
from starlette.types import ASGIApp

from market.web.di import bootstrap
from market.web.routes import routes

bootstrap()

def create_app() -> ASGIApp:
    return Starlette(
        debug=True,
        routes=routes,
    )
