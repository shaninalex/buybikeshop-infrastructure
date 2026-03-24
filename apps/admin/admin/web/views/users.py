from starlette.endpoints import HTTPEndpoint

from admin.core.service_kratos import ServiceKratos
from admin.web.di import resolve
from admin.web.templating import Templates


class UsersPage(HTTPEndpoint):
    async def get(self, request):
        templates = resolve(Templates)
        sk = resolve(ServiceKratos)
        return templates.TemplateResponse(request, "views/users.html", {
            "users": sk.get_users()
        })