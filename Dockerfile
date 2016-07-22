FROM alpine:latest

MAINTAINER Cesar Gutierrez Tineo <cesar@tineo.mobi>

WORKDIR "/opt"

ADD .docker_build/homepage /opt/bin/homepage
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/homepage"]

