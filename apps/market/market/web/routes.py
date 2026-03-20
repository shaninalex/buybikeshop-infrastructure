import os

from starlette.routing import Route, Mount
from starlette.staticfiles import StaticFiles

from market.web.views import HomePage
from market.web.views.product import ProductPage

routes = [
    Mount("/static", app=StaticFiles(directory=os.environ.get("APP_MARKET_WEB_STATIC_PATH")), name="static"),
    Route("/", HomePage, name="home"),
    Route("/product/{product_id:int}", ProductPage, name="product"),
]
