import os

from starlette.routing import Route, Mount
from starlette.staticfiles import StaticFiles

from admin.web.views import HomePage

routes = [
    Mount("/static", app=StaticFiles(directory=os.environ.get("APP_ADMIN_WEB_STATIC_PATH")), name="static"),
    Route("/", HomePage, name="home"),
]
