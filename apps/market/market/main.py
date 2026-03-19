"""
Runs webserver
"""
import os

import uvicorn
from dotenv import load_dotenv

from web.app import create_app

app = create_app()

if __name__ == "__main__":
    load_dotenv()
    port = int(os.getenv("APP_MARKET_PORT", 80))
    uvicorn.run(
        app="run.web:app",
        port=port,
    )
