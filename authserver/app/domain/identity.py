from dataclasses import dataclass
from datetime import datetime
from uuid import UUID


@dataclass()
class Identity:
    id: UUID
    full_name: str
    email: str
    active: bool
    created_at: datetime
