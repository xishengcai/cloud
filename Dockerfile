FROM alpine:latest

MAINTAINER xishengcai <cc710917049@163.com>

COPY ./bin/cloud /usr/local/bin
COPY ./conf  /opt/conf
COPY ./template /opt/template
COPY ./docs /opt/docs
COPY ./image_ftp /opt/image_ftp

RUN chmod +x /usr/local/bin/cloud

WORKDIR /opt

EXPOSE 80

CMD ["cloud"]
