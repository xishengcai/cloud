#!/usr/bin/env bash
# made by Caixisheng  Fri Nov 9 CST 2018

#chec user
[[ $UID -ne 0 ]] && { echo "Must run in root user !";exit; }

isOuter=true

checkNet(){
  # -c: 表示次数，1 为1次
  # -w: 表示deadline, time out的时间，单位为秒，100为100秒。
  ping -c 1 -w 3 www.google.com
  if [[ $? != 0 ]];then
    echo "in inter"
    isOuter=false
  else
    echo "in out"
  fi
}

checkNet

set -e

rm -rf /etc/cni/net.d
rm -rf /etc/kubernetes
rm -rf /var/lib/etcd
kubeadm reset --force

# https://pkg.go.dev/k8s.io/kubernetes@v1.17.2/cmd/kubeadm/app/apis/kubeadm/v1beta2
cat > kubeadm-config.yaml <<EOF
apiVersion: kubeadm.k8s.io/v1beta2
controlPlaneEndpoint: {{.ControlPlaneEndpoint}}
imageRepository: {{.Registry}}
kind: ClusterConfiguration
kubernetesVersion: {{.Version}}
controllerManager:
  extraArgs:
    experimental-cluster-signing-duration: 87600h0m0s
networking:
  podSubnet: {{.PodCidr}}
  serviceSubnet: {{.ServiceCidr}}
---
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
mode: ipvs
---
kind: KubeletConfiguration
apiVersion: kubelet.config.k8s.io/v1beta1
cgroupDriver: systemd
#---
#apiVersion: kubelet.config.k8s.io/v1beta1
#kind: KubeletConfiguration
#evictionHard:  # 配置硬驱逐阈值
#  memory.available: "500Mi"
#  cpu.available: 300m
EOF

echo "kubeadm init"
kubeadm init --config=/root/kubeadm-config.yaml --upload-certs

# To start using your cluster, you need to run the following as a regular user:
mkdir -p /root/.kube
\cp /etc/kubernetes/admin.conf /root/.kube/config


# install helm3
if $isOuter; then
#  curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
#  chmod 700 get_helm.sh
#  ./get_helm.sh

  curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
else
  curl -fsSl -o helm-canary-linux-amd64.tar.gz https://soft-package-xisheng.oss-cn-hangzhou.aliyuncs.com/k8s/helm-v3.9.3-linux-amd64.tar.gz
  tar xzvf helm-canary-linux-amd64.tar.gz
  mv linux-amd64/helm /usr/local/bin/
fi

kubectl taint  nodes --all node-role.kubernetes.io/master-
