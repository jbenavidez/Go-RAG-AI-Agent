FROM alpine:latest

RUN mkdir /app

COPY CliService /app

CMD [ "/app/CliService"]