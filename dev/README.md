# DEV

Infrastructure for local development.

Key feature that most services should be running locally because of security, redirects, cookies etc.
Need to install localy:
- [kratos](https://www.ory.com/docs/identities)
- [hydra](https://www.ory.com/docs/oauth2-oidc)
- [keto](https://www.ory.com/docs/keto)
- [oathkeeper](https://www.ory.com/docs/oathkeeper)

Ory-independent services in via docker compose:
- postgres
- rabbitmq
- redis

> List will be changed during development...
---

## Install

Navigate to `~/.local/bin` directory, then install:

**Kratos**

```bash
bash <(curl https://raw.githubusercontent.com/ory/meta/master/install.sh) -d -b . kratos v25.4.0
./kratos help
```

**Hydra**

```bash
bash <(curl https://raw.githubusercontent.com/ory/meta/master/install.sh) -d -b . hydra v25.4.0
./hydra help
```

**Keto**

```bash
bash <(curl https://raw.githubusercontent.com/ory/meta/master/install.sh) -d -b . keto v25.4.0
./keto help
```

**Oathkeeper**

```bash
bash <(curl https://raw.githubusercontent.com/ory/meta/master/install.sh) -d -b . oathkeeper v25.4.0
./oathkeeper help
```

---

## Start

```bash
$ ./scripts/start-dev-server.sh
```
