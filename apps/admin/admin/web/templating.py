import datetime
import os

import jinja2
from starlette.templating import Jinja2Templates


class Templates(Jinja2Templates):
    def __init__(self) -> None:
        super().__init__(directory=os.environ.get("APP_ADMIN_WEB_TEMPLATES_PATH"))

        # example setting different pipes
        self.env.filters["dateformat"] = _pipe_dateformat

    def from_string(self, source: str) -> jinja2.Template:
        return self.env.from_string(source)


def _pipe_dateformat(value: datetime.date | str) -> str:
    """
    Format date time in standard format.
    Usage:
        {{ post.date_published | dateformat }}
    :param value: datetime
    :return: string
    """
    if isinstance(value, str):
        value = datetime.date.fromisoformat(value)
    return value.strftime("%b %d, %Y")
