#!/bin/bash


IMAGE_NAME="devops-dashboard"
VERSION="latest"
REGISTRY="xishengcai"

cat <<EOF > Dockerfile
FROM nginx
COPY nginx.conf /etc/nginx/
COPY dist/     /usr/share/nginx/html
EOF

docker build -t ${REGISTRY}/${IMAGE_NAME}:${VERSION} ./
docker push ${REGISTRY}/${IMAGE_NAME}:${VERSION}
docker rmi ${REGISTRY}/${IMAGE_NAME}:${VERSION}
