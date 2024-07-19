FROM debian


WORKDIR /microService

COPY microService microService
COPY conf.toml conf.toml

EXPOSE 8000


RUN apt update -y && apt upgrade -y && apt install iproute2 iputils-ping postgresql-client -y

CMD [ "./microService"]
