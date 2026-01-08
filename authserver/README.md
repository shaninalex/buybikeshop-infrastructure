# AuthServer

```bash

# Run server
python authserver/main.py serve run --config=./config/config.dev.yaml

# Create identity
python authserver/main.py identities create \
  --fullname="Full Name" \
  --email=test@test.com \
  --config=./config/config.dev.yaml

# Delete identity
python authserver/main.py identities delete \
  --id=12 \
  --config=./config/config.dev.yaml
```
