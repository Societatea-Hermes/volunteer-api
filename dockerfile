FROM alpine:latest

RUN mkdir /app

COPY hermesAPI /app

CMD [ "/app/hermesAPI"]