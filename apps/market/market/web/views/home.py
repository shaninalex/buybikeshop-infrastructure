from starlette.endpoints import HTTPEndpoint

from catalog.product_pb2 import ProductListRequest
from market.core.connector import CatalogConnector
from market.web.di import resolve
from market.web.templating import Templates


class HomePage(HTTPEndpoint):
    async def get(self, request):
        templates = resolve(Templates)
        connector: CatalogConnector = resolve(CatalogConnector)
        try:
            resp = connector.product_list(ProductListRequest())
            return templates.TemplateResponse(request, "views/home.html", {
                "products": resp.products
            })
        except Exception as e:
            return templates.TemplateResponse(request, "views/500.html", {
                "error": f"{e}"
            })
