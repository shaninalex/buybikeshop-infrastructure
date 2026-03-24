# Admin

```bash
uv run --project apps/admin uvicorn admin.main:app \
    --reload \
    --env-file apps/admin/.env \
    --port 8003
```