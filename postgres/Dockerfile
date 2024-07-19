FROM postgres:latest
ENV POSTGRES_PASSWORD admin

COPY postgresql.conf /etc/postgresql.conf

COPY init.sql /docker-entrypoint-initdb.d/

RUN chown postgres:postgres /etc/postgresql.conf

RUN apt update -y && apt upgrade -y && apt install iproute2 iputils-ping postgresql-client -y
