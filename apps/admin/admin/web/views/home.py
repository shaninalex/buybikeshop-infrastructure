from starlette.endpoints import HTTPEndpoint

from admin.web.di import resolve
from admin.web.templating import Templates


class HomePage(HTTPEndpoint):
    async def get(self, request):
        templates = resolve(Templates)
        return templates.TemplateResponse(request, "views/home.html", {})