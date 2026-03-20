FROM postgres:16

WORKDIR /root/

COPY . .

#COPY ./migrations ./migrations
#COPY ./bin ./bin
#COPY ./run.sh ./run.sh

RUN chmod +x /root/migrate.sh
