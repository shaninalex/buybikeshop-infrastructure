# Seeder

```bash
# generate products
$ uv run --package seeder python -m seeder.main \
    --config ./resources/database/seeder/config.yaml \
    seed-products

# generate customers
$ uv run --package seeder python -m seeder.main \
    --config ./resources/database/seeder/config.yaml \
    seed-customers
    
# generate orders
$ uv run --package seeder python -m seeder.main \
    --config ./resources/database/seeder/config.yaml \
    seed-orders
```
