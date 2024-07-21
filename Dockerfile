# GO SERVER
FROM debian

WORKDIR /microServiceDir

# COPY test/kafka kafka

COPY conf.toml conf.toml
COPY microService microService

EXPOSE 8000

RUN apt update -y && apt upgrade -y && apt install iproute2 iputils-ping postgresql-client -y

CMD [ "./microService"]