#!/usr/bin/env bash
# shellcheck disable=SC2006
export VERSION=`docker version  --format="{{.Server.Version}}"`
if [[ $VERSION == "20.10.24" ]];then
  echo "docker version 20.10.24 already install"
  exit 0
fi

set -e
echo "clean env"
yum remove -y docker docker-common container-selinux docker-selinux docker-engine
rm -rf /var/lib/docker

echo "install docker 20.10.24"
yum install -y yum-utils

yum-config-manager \
    --add-repo \
    https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

yum clean packages
#查看docker-ce版本并且安装
yum list docker-ce --showduplicates | sort -r
yum install -y docker-ce-cli-20.10.24 docker-ce-20.10.24 containerd.io docker-compose-plugin docker-ce-rootless-extras-20.10.24


echo "config docker daemon"
mkdir -p /etc/docker
cat > /etc/docker/daemon.json <<EOF
{
  "data-root": "/var/lib/docker",
  "storage-driver": "overlay2",
  "exec-opts": [
    "native.cgroupdriver=systemd",
    "overlay2.override_kernel_check=true"
  ],
  "live-restore": true,
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "registry-mirrors": ["https://wms804s3.mirror.aliyuncs.com"]
}
EOF

systemctl enable docker.service
systemctl daemon-reload
systemctl enable docker
systemctl restart docker

docker info
