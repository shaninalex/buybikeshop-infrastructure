from starlette.endpoints import HTTPEndpoint

from market.core.connector import CatalogConnector
from market.web.di import resolve
from market.web.templating import Templates


class ProductPage(HTTPEndpoint):
    async def get(self, request):
        templates = resolve(Templates)
        product_id = int(request.path_params['product_id'])
        connector: CatalogConnector = resolve(CatalogConnector)
        try:
            resp = connector.product_get(product_id)
            return templates.TemplateResponse(request, "views/product.html", {
                "product": resp.product
            })
        except Exception as e:
            return templates.TemplateResponse(request, "views/500.html", {
                "error": f"{e}"
            })
