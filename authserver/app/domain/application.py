from dataclasses import dataclass
from datetime import datetime
from typing import List
from uuid import UUID


@dataclass()
class ApplicationRedirectUrl:
    id: UUID
    application_id: UUID
    redirect_url: str
    active: bool


@dataclass()
class Application:
    id: UUID
    app_name: str
    active: bool
    client_id: str
    client_secret: str
    url_home_page: str
    created_at: datetime

    redirect_urls: List[ApplicationRedirectUrl]
