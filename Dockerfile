FROM scratch

MAINTAINER Indra Octama omyank2007i@gmail.com

ADD main ./

ARG appname=dockerinaja
ARG http_port=1323

ENTRYPOINT ["/main"]

ENV PORT $http_port
ENV DB_HOST dev.indraoctama.com
ENV DB_PORT 3306
ENV DB_USER developer
ENV DB_PASS dev123

EXPOSE $http_port
