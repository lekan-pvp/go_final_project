FROM ubuntu:latest

WORKDIR /usr/local/app

ENV TODO_PORT=7540
EXPOSE ${TODO_PORT}

COPY ./web ./web
COPY ./go1f .

CMD ["./go1f"]