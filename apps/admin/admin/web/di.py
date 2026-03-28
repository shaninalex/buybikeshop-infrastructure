import os

from admin.core.service_kratos import ServiceKratos
from admin.web.templating import Templates
from common_tools.di import Container


def _configure(container: Container) -> None:
    container.register(ServiceKratos, instance=ServiceKratos(os.environ.get("APP_ADMIN_KRATOS_ADMIN")))
    container.register(Templates, instance=Templates())


def create_container() -> Container:
    return Container(_configure)


_container = create_container()
bootstrap = _container.bootstrap
resolve = _container.resolve
