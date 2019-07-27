FROM golang:alpine

MAINTAINER Indra Octama omyank2007i@gmail.com

ADD . /go/src/dockerinaja

ARG appname=e-document-api
ARG http_port=1323

RUN apk update && apk add git

RUN go get -d github.com/go-sql-driver/mysql
RUN go get -d github.com/labstack/echo
RUN go install dockerinaja

ENTRYPOINT /go/bin/dockerinaja

ENV PORT $http_port
ENV DB_HOST dev.indraoctama.com
ENV DB_PORT 3306
ENV DB_USER developer
ENV DB_PASS dev123

EXPOSE $http_port

RUN mkdir -p /go/src/dockerinaja
COPY . /go/src/dockerinaja
WORKDIR /go/src/dockerinaja

CMD go run main.go