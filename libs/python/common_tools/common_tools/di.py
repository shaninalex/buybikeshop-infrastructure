from typing import Callable, Type, TypeVar, Union, Any

import punq

T = TypeVar("T")
K = TypeVar("K", bound=Union[Type[Any], str])


class Container:
    def __init__(self, configure: Callable[["Container"], None]):
        self._impl = punq.Container()
        self._configure = configure
        self._is_configured = False

    def register(self, key: K, *, instance: Any) -> None:
        """Register either a type or a string key"""
        self._impl.register(key, instance=instance)

    def bootstrap(self) -> None:
        if self._is_configured:
            return  # pragma: no cover

        self._configure(self)
        self._is_configured = True

    def resolve(self, key: K) -> Any:
        assert self._is_configured, "DI not configured: call bootstrap()"
        return self._impl.resolve(key)
