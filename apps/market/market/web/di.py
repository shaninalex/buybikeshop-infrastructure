"""Web app DI"""

from market.core.di import Container
from market.web.templating import Templates


def _configure(container: Container) -> None:
    container.register(Templates, instance=Templates())


def create_container() -> Container:
    return Container(_configure)


_container = create_container()
bootstrap = _container.bootstrap
resolve = _container.resolve
