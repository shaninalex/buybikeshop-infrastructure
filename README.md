## Buybikeshop

> It's a playground, a place to test ideas and practice.

Organization infrastructure - db, apps, servers, ui, etc...

### Run:

```bash
# services configs
cp apps/admin/.env.example apps/admin/.env
cp apps/datasource/config/config.example.yml apps/datasource/config/config.yml
cp apps/market/.env.example apps/market/.env
cp apps/office/config/config.example.yml apps/office/config/config.yml
cp apps/warehouse/config/config.example.yml apps/warehouse/config/config.yml
# modify copied configs with proper variables

./run.sh generate_grpc
./run.sh start

# TODO: create users
# ( currently I use Postman, see docs/ory_hydra-ory_kratos_api_integration.json)

# generate mock data
./run.sh seed
```
