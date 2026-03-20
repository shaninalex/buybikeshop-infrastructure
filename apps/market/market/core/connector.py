import grpc.aio

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
from catalog_pb2_grpc import CatalogStub


class CatalogConnector:
    def __init__(self, channel: grpc.aio.Channel):
        self._stub = CatalogStub(channel)

    def product_list(self, request: ProductListRequest) -> ProductListReply:
        return self._stub.ProductList(request)

    def product_get(self, product_id: int) -> ProductGetReply:
        return self._stub.ProductGet(ProductGetRequest(id=product_id))

    def product_variant_list(
        self, request: ProductVariantListRequest
    ) -> ProductVariantListReply:
        return self._stub.ProductVariantList(request)

    def product_variant_get(self, variant_id: int) -> ProductVariantGetReply:
        return self._stub.ProductVariantGet(
            ProductVariantGetRequest(id=variant_id)
        )
