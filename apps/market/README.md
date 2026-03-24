```bash
uv run --project apps/market uvicorn market.main:app \
  --reload \
  --env-file apps/market/.env \
  --port 8004
```
