from starlette.applications import Starlette
from starlette.types import ASGIApp

from admin.web.di import bootstrap
from admin.web.routes import routes

bootstrap()

def create_app() -> ASGIApp:
    return Starlette(
        debug=True,
        routes=routes,
    )
