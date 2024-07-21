# MSG HANDLER
FROM debian

WORKDIR /msgHandlerDir


COPY msgHandler msgHandler

RUN apt update -y && apt upgrade -y && apt install iproute2 iputils-ping postgresql-client -y

CMD [ "./msgHandler"]