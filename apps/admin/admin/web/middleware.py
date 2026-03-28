from starlette.requests import Request
from starlette.types import ASGIApp, Receive, Scope, Send

# NOTE: this is not very beautiful or elegant... Better will be using separate admin-frontend-client and make
#       proper API requests with all possible security measures from Oathkeeper, Hydra, Kratos etc...

class CSRFCookieToHeaderMiddleware:
    def __init__(self, app: ASGIApp, cookie_name: str = "csrftoken", header_name: str = "x-csrftoken"):
        self.app = app
        self.cookie_name = cookie_name
        self.header_name = header_name

    async def __call__(self, scope: Scope, receive: Receive, send: Send):
        if scope["type"] == "http":
            request = Request(scope)
            token = request.cookies.get(self.cookie_name)
            if token and self.header_name not in request.headers:
                scope["headers"].append((self.header_name.encode(), token.encode()))

        await self.app(scope, receive, send)
