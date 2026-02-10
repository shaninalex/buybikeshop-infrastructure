from datetime import datetime

from sqlalchemy import (
    UUID,
    Boolean,
    Column,
    DateTime,
    ForeignKey,
    String,
    Table,
)

from . import metadata


def default_current_date():
    return datetime.now()


identities = Table(
    "identities",
    metadata,
    Column("id", UUID, primary_key=True),
    Column("fullname", String(30)),
    Column("email", String(80)),
    Column("active", Boolean(False)),
    Column("created_at", DateTime),
)

applications = Table(
    "applications",
    metadata,
    Column("id", UUID, primary_key=True),
    Column("app_name", String(30)),
    Column("active", Boolean(False)),
    Column("client_id", String()),
    Column("client_secret", String()),
    Column("home_url", String()),
    Column("created_at", DateTime, default=default_current_date),
)


application_redirect_uris = Table(
    "application_redirect_uris",
    metadata,
    Column("id", UUID, primary_key=True),
    Column("application_id", UUID, ForeignKey("applications.id")),
    Column("redirect_uri", String()),
    Column("active", Boolean(False)),
)

credentials = Table(
    "credentials",
    metadata,
    Column("id", UUID, primary_key=True),
    Column("identity_id", UUID, ForeignKey("identities.id")),
    Column("provider", String()),
    Column("provider_user_id", String()),
    Column("email", String()),
    Column("pwd_hash", String()),
    Column("created_at", DateTime, default=default_current_date),
)
