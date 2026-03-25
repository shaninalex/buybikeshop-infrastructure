from starlette.endpoints import HTTPEndpoint

from admin.core.service_kratos import ServiceKratos
from admin.web.di import resolve
from admin.web.templating import Templates


class UserDetailPage(HTTPEndpoint):
    async def get(self, request):
        templates = resolve(Templates)
        sk: ServiceKratos = resolve(ServiceKratos)
        user = sk.get_identity(request.path_params.get('id'))
        return templates.TemplateResponse(request, "views/user-detail.html", {
            "user": user,
        })

    async def post(self, request):
        templates = resolve(Templates)
        sk: ServiceKratos = resolve(ServiceKratos)
        user = sk.get_identity(request.path_params.get('id'))
        return templates.TemplateResponse(request, "views/user-detail.html", {
            "user": user,
        })
