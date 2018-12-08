FROM golang:1.11.1

MAINTAINER Indra Octama indra.octama@orori.com

ADD . /go/src/dockerinaja

RUN go get -d github.com/go-sql-driver/mysql
RUN go get -d github.com/labstack/echo
RUN go install dockerinaja

ENTRYPOINT /go/bin/dockerinaja

EXPOSE 1323

RUN mkdir -p /go/src/dockerinaja
COPY . /go/src/dockerinaja
WORKDIR /go/src/dockerinaja

CMD go run main.go