#!/bin/bash

#export GOPROXY=https://goproxy.cn
export GO111MODULE=on

IMAGE_NAME="cloud"
VERSION="release-V1"
REGISTRY="registry.cn-hangzhou.aliyuncs.com/xisheng"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/cloud ./main.go

if [ $? -ne 0 ]; then
    echo "build ERROR"
    exit 1
fi

echo build success

cat <<EOF > Dockerfile
FROM alpine:latest

MAINTAINER xishengcai <cc710917049@163.com>

COPY ./bin/cloud /usr/local/bin
COPY ./conf  /opt/conf
COPY ./template /opt/template
COPY ./docs /opt/docs
#COPY ./image_ftp /opt/image_ftp

RUN chmod +x /usr/local/bin/cloud

WORKDIR /opt

EXPOSE 80

CMD ["cloud"]
EOF

docker build -t ${REGISTRY}/${IMAGE_NAME}:${VERSION} ./
docker push ${REGISTRY}/${IMAGE_NAME}:${VERSION}
docker rmi ${REGISTRY}/${IMAGE_NAME}:${VERSION}
