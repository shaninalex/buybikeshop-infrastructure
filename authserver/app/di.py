from typing import Callable, Type, TypeVar, cast

import punq
from app.db import database
from databases import Database

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
    # container.register(Templates, instance=Templates())
    container.register(Database, instance=database)
    # categories_repository = CategoriesRepository(database)
    # products_variants_repository = ProductsVariantsRepository(database)
    # products_repository = ProductsRepository(database, products_variants_repository)
    # container.register(CategoriesRepository, instance=categories_repository)
    # container.register(
    #     ProductsVariantsRepository, instance=products_variants_repository
    # )
    # container.register(ProductsRepository, instance=products_repository)
    # web_entities = EntityWebService(
    #     products_variants_repository, categories_repository, products_repository
    # )
    # container.register(EntityWebService, instance=web_entities)
    pass


def create_container() -> Container:
    return Container(_configure)


_container = create_container()
bootstrap = _container.bootstrap
resolve = _container.resolve
