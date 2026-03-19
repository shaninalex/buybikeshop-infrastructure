# import grpc
from starlette.applications import Starlette
from starlette.types import ASGIApp

# from market.core.connector import CatalogConnector

def create_app() -> ASGIApp:
    # with grpc.insecure_channel('localhost:50051') as channel:
    #     conn = CatalogConnector(channel)
    #     result = await conn.product_get(1)
    #     result.product.title

    return Starlette(
        debug=True,
    )
