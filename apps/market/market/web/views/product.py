from starlette.endpoints import HTTPEndpoint

from market.web.di import resolve
from market.web.templating import Templates


class ProductPage(HTTPEndpoint):
    async def get(self, request):
        templates = resolve(Templates)
        return templates.TemplateResponse(request, "views/product.html", {})
