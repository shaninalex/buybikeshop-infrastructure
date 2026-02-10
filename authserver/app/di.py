from typing import Callable, Type, TypeVar, cast

import punq

# from app.db import database
from app.repositories import ApplicationRepository

# from databases import Database

T = TypeVar("T")


class Container:
    def __init__(self, configure: Callable[["Container"], None]):
        self._impl = punq.Container()
        self._configure = configure
        self._is_configured = False

    def register(self, type_: Type[T], *, instance: T) -> None:
        self._impl.register(type_, instance=instance)

    def bootstrap(self) -> None:
        if self._is_configured:
            return  # pragma: no cover

        self._configure(self)
        self._is_configured = True

    def resolve(self, type_: Type[T]) -> T:
        assert self._is_configured, "DI not configured: call bootstrap()"
        return cast(T, self._impl.resolve(type_))


def _configure(container: Container) -> None:
    # container.register(Database, instance=database)
    # application_repository = ApplicationRepository(database)
    # container.register(ApplicationRepository, instance=application_repository)
    pass


def create_container() -> Container:
    print("Create DI container")
    return Container(_configure)


_container = create_container()
bootstrap = _container.bootstrap
resolve = _container.resolve
