"""Web app DI"""

from common_tools.di import Container

from admin.web.templating import Templates


def _configure(container: Container) -> None:
    container.register(Templates, instance=Templates())


def create_container() -> Container:
    return Container(_configure)


_container = create_container()
bootstrap = _container.bootstrap
resolve = _container.resolve
