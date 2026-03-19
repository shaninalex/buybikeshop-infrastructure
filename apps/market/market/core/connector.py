import grpc.aio

from catalog_pb2_grpc import CatalogStub
from catalog.product_pb2 import (
    ProductGetReply,
    ProductGetRequest,
    ProductListReply,
    ProductListRequest,
    ProductVariantGetReply,
    ProductVariantGetRequest,
    ProductVariantListReply,
    ProductVariantListRequest,
)


class CatalogConnector:
    def __init__(self, channel: grpc.aio.Channel):
        self._stub = CatalogStub(channel)

    async def product_list(self, request: ProductListRequest) -> ProductListReply:
        return await self._stub.ProductList(request)

    async def product_get(self, product_id: int) -> ProductGetReply:
        return await self._stub.ProductGet(ProductGetRequest(id=product_id))

    async def product_variant_list(
        self, request: ProductVariantListRequest
    ) -> ProductVariantListReply:
        return await self._stub.ProductVariantList(request)

    async def product_variant_get(self, variant_id: int) -> ProductVariantGetReply:
        return await self._stub.ProductVariantGet(
            ProductVariantGetRequest(id=variant_id)
        )
