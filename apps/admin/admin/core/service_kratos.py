from typing import List

import ory_kratos_client
from ory_kratos_client import Identity


class ServiceKratos:
    def __init__(self, url: str):
        self.configuration = ory_kratos_client.Configuration(
            host=url
        )
        self.root = url

    def get_users(self) -> List[Identity]:
        with ory_kratos_client.ApiClient(self.configuration) as api_client:
            api_instance = ory_kratos_client.IdentityApi(api_client)
            per_page = 25
            page = 1
            page_size = 25
            resp = api_instance.list_identities(per_page=per_page, page=page, page_size=page_size)
            return resp
