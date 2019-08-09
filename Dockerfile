FROM scratch

MAINTAINER Indra Octama omyank2007i@gmail.com

ADD main ./

ARG appname=dockerinaja
ARG http_port=1323

ENTRYPOINT ["/main"]

EXPOSE $http_port