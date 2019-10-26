FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --no-cache --update add ca-certificates tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 8080

COPY my-server-go /app

CMD /app/my-server-go