# Using Multi Stage Docker

# GOLANG BUILDER
FROM golang:1.13.3-alpine3.10 as builder

RUN apk update && apk upgrade && \
    apk --no-cache --update add git make

WORKDIR /go/src/github.com/andikanugraha11/golang-rest-docker-kube

COPY . .

RUN go build -o my-go-server main.go

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --no-cache --update add ca-certificates tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 8080

COPY --from=builder /go/src/github.com/andikanugraha11/golang-rest-docker-kube /app

CMD /app/my-go-server