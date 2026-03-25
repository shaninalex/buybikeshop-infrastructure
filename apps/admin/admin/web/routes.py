import os

from starlette.routing import Route, Mount
from starlette.staticfiles import StaticFiles

from admin.web.views import HomePage
from admin.web.views.user_detail import UserDetailPage
from admin.web.views.users import UsersPage

routes = [
    Mount("/static", app=StaticFiles(directory=os.environ.get("APP_ADMIN_WEB_STATIC_PATH")), name="static"),
    Route("/", HomePage, name="home"),
    Route("/users", UsersPage, name="users"),
    Route("/users/{id:str}", UserDetailPage, name="user-detail", methods=["GET", "POST"]),
]
