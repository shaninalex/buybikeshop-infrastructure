from starlette.endpoints import HTTPEndpoint

from market.web.di import resolve
from market.web.templating import Templates


class HomePage(HTTPEndpoint):
    async def get(self, request):
        templates = resolve(Templates)
        products = [
            {
                "id": 1,
                "title": "product A",
                "summary": "Cards support a wide variety of content, including images, text, list groups, links, and more.",
                "image": "images/2023-Author-Aura-XR2-768x448-2150028769.jpg",
            },
            {
                "id": 2,
                "title": "product B",
                "summary": "Cards support a wide variety of content, including images, text, list groups, links, and more.",
                "image": "images/2023-Author-Compact-768x496-484620391.jpg"
            },
            {
                "id": 4,
                "title": "product C",
                "summary": "Cards support a wide variety of content, including images, text, list groups, links, and more.",
                "image": "images/ua42838601_author_bicycle_charisma_55_2015-2375528177.jpg"
            }
        ]
        return templates.TemplateResponse(request, "views/home.html", {
            "products": products,
        })
