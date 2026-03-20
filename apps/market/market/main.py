"""
Runs webserver
"""
import os

import uvicorn

from market.web.app import create_app

app = create_app()

if __name__ == "__main__":
    port = int(os.getenv("APP_MARKET_PORT"))
    config = uvicorn.Config("run.web:app", port=port, log_level="info")
    server = uvicorn.Server(config)
    server.run()
