import os

from starlette.applications import Starlette
from starlette.middleware import Middleware
from starlette.types import ASGIApp
from starlette_csrf import CSRFMiddleware

from admin.web.di import bootstrap
from admin.web.middleware import CSRFCookieToHeaderMiddleware
from admin.web.routes import routes

bootstrap()


def create_app() -> ASGIApp:
    return Starlette(
        debug=True,
        routes=routes,
        middleware=[
            Middleware(CSRFCookieToHeaderMiddleware),
            Middleware(CSRFMiddleware, secret=os.environ.get("APP_ADMIN_SECRET")),
        ]
    )
