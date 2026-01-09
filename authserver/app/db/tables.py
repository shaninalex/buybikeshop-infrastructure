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
    Column("created_at", DateTime),
)


application_redirect_uris = Table(
    "application_redirect_uris",
    metadata,
    Column("id", UUID, primary_key=True),
    Column("application_id", UUID, ForeignKey("applications.id")),
    Column("redirect_uri", String()),
    Column("active", Boolean(False)),
)
