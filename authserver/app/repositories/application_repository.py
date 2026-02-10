from typing import List
from uuid import UUID

from app.db.tables import application_redirect_uris, applications
from app.domain import Application, ApplicationRedirectUrl

# from databases import Database
from sqlalchemy import insert, outerjoin, select


class ApplicationRepository:
    def __init__(self, db):
        self._db = db

    async def list(self) -> List[Application]:
        q = select(applications, application_redirect_uris).select_from(
            outerjoin(
                applications,
                application_redirect_uris,
                applications.c.id == application_redirect_uris.c.application_id,
            )
        )
        result = await self._db.fetch_all(q)
        apps: dict[UUID, Application] = {}

        for row in result:
            print(row._mapping)
            app_row = row._mapping[applications]
            redirect_row = row._mapping[application_redirect_uris]

            app_id = app_row.id

            if app_id not in apps:
                apps[app_id] = Application(
                    id=app_row.id,
                    app_name=app_row.app_name,
                    active=app_row.active,
                    client_id=app_row.client_id,
                    client_secret=app_row.client_secret,
                    url_home_page=app_row.home_url,
                    created_at=app_row.created_at,
                    redirect_urls=[],
                )

            if redirect_row.id is not None:
                apps[app_id].redirect_urls.append(
                    ApplicationRedirectUrl(
                        id=redirect_row.id,
                        application_id=redirect_row.application_id,
                        redirect_uri=redirect_row.redirect_uri,
                        active=redirect_row.active,
                    )
                )

        return list(apps.values())

    async def save(self, app: Application):
        await self._db.execute(
            insert(applications).values(
                id=app.id,
                app_name=app.app_name,
                active=app.active,
                client_id=app.client_id,
                client_secret=app.client_secret,
                home_url=app.url_home_page,
                created_at=app.created_at,
            )
        )
        if app.redirect_urls:
            await self._db.execute(
                insert(application_redirect_uris).values(
                    [
                        {
                            "id": r.id,
                            "application_id": r.application_id,
                            "redirect_uri": r.redirect_uri,
                            "active": r.active,
                        }
                        for r in app.redirect_urls
                    ]
                )
            )
