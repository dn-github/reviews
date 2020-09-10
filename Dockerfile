FROM alpine:3.4

COPY reviews .
ENTRYPOINT ./reviews
EXPOSE 3001