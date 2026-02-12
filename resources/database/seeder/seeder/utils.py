import random

DEFAULT_MEDIA = [
    "/img/bike0.jpg",
    "/img/bike1.jpg",
    "/img/bike2.jpg",
    "/img/bike3.jpg",
    "/img/bike4.jpg"
]


def get_random_media() -> List[str]:
    count = random.randint(0, len(DEFAULT_MEDIA))
    return random.sample(DEFAULT_MEDIA, count)
