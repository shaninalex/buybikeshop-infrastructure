FROM postgres:16

ARG MIGRATE_VERSION=v4.18.3

WORKDIR /root/

RUN apt-get update && apt-get install -y curl \
    && curl -fsSL https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz \
       | tar -xz -C /root/ \
    && apt-get purge -y curl && apt-get autoremove -y && rm -rf /var/lib/apt/lists/*

COPY . .

RUN chmod +x /root/migrate.sh
