# AuthServer

```bash
# Run server
python authserver/main.py \
    --config=./authserver/config/config.dev.yaml \
    serve run

# Create identity
python authserver/main.py identities create \
    --config=./authserver/config/config.dev.yaml \
    --fullname="Full Name" \
    --email=test@test.com \

# Delete identity
python authserver/main.py \
    --config=./authserver/config/config.dev.yaml \
    identities delete \
    --id=12
  
# Application create
python authserver/main.py \
    --config=./authserver/config/config.dev.yaml \
    applications create \
    --json=./authserver/resources/application_payload.json
```
