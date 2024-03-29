#!/usr/bin/env bash
# made by Caixisheng  Fri Nov 9 CST 2018

#chec user
[[ $UID -ne 0 ]] && { echo "Must run in root user !";exit; }

export VERSION=`kubeadm version -o=short`
if [[ $VERSION == "v{{.Version}}" ]];then
  echo "kubeadm already install"
  set -e
  rm -rf /etc/cni/net.d
  rm -rf /etc/kubernetes
  rm -rf /var/lib/etcd
  kubeadm reset --force
  exit 0
fi
set -e

echo "添加kubernetes国内yum源"
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
       http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

# Set SELinux in permissive mode (effectively disabling it)
cat <<EOF >  /etc/sysctl.d/k8s.conf
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
vm.swappiness=0
EOF

sysctl --system
swapoff -a
sed -ri 's/.*swap.*/#&/' /etc/fstab

cat <<EOF > /etc/sysconfig/modules/ipvs.modules
#!/bin/bash
modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack_ipv4
EOF

chmod 755 /etc/sysconfig/modules/ipvs.modules && bash /etc/sysconfig/modules/ipvs.modules && lsmod | grep -e ip_vs -e nf_conntrack_ipv4
yum install ipset ipvsadm -y


# if this node is ready, you want reJoin, maybe occure device busy
set +e
kubeadm reset --force

set -e
rm -rf /var/lib/cni/
rm -rf /var/lib/etcd/
rm -rf /etc/kubernetes
rm -rf /var/lib/kubelet
rm -rf /var/lib/dockershim
rm -rf /etc/cni/net.d

yum -y  remove kubeadm kubectl kubelet
yum -y install kubelet-{{.Version}} kubeadm-{{.Version}} kubectl-{{.Version}} --setopt=obsoletes=0

systemctl daemon-reload
systemctl enable kubelet
systemctl start kubelet

#yum install wget -y
#wget https://soft-package-xisheng.oss-cn-hangzhou.aliyuncs.com/k8s/kubeadm-1.17.2
#/bin/cp -rf kubeadm-1.17.2 /usr/local/bin/kubeadm
kubeadm reset --force

## 关闭防火墙
systemctl disable firewalld
systemctl stop firewalld
