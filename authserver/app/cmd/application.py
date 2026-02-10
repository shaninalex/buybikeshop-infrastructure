import json
from datetime import datetime
from typing import List
from uuid import uuid4

from app import domain
from app.config import Config

# from app.db import database
from app.di import bootstrap, resolve
from app.repositories import ApplicationRepository

bootstrap()


def parse_application_from_file(path: str) -> domain.Application:
    with open(path, "r") as f:
        data = json.load(f)
        application_id = uuid4()
        redirect_urls: List[domain.ApplicationRedirectUrl] = []
        for u in data["redirect_urls"]:
            redirect_urls.append(
                domain.ApplicationRedirectUrl(
                    id=uuid4(),
                    application_id=application_id,
                    redirect_uri=u["redirect_uri"],
                    active=True,
                )
            )

        app: domain.Application = domain.Application(
            id=application_id,
            app_name=data["app_name"],
            active=False,
            client_id=data["client_id"],
            client_secret=data["client_secret"],
            url_home_page=data["url_home_page"],
            created_at=datetime.now(),
            redirect_urls=redirect_urls,
        )
        return app


async def application_create(payload_path: str, config: Config):
    # await database.connect()
    print("Create application from file: ", payload_path)

    app = parse_application_from_file(payload_path)

    repository = resolve(ApplicationRepository)

    result = await repository.list()

    for a in result:
        if a.app_name == app.app_name:
            raise Exception("alredy exists", app.app_name)

    await repository.save(app)
    # await database.disconnect()
