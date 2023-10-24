#!/bin/bash

#Requirements:
#
#Kubernetes must be configured to use CNI (see Network Plugin Requirements)
#Linux kernel >= 4.9.17
#内核升级指南：https://www.cnblogs.com/ding2016/p/10429640.html

set -e

kernel=`uname -r | cut -d "-" -f 1`

kernel_v1=`echo $kernel | cut -d "." -f 1`
kernel_v2=`echo $kernel | cut -d "." -f 2`
kernel_v3=`echo $kernel | cut -d "." -f 3`

echo `uname -r`
if [[ $kernel_v1 -lt 4 ]];then
  echo "$kernel 版本低于 4.9.17"
  sh upgrade_kernel.sh
  exit 1
fi

if [[ $kernel_v1 -lt 4 ]] && [[ $kernel_v2 -lt 9 ]];then
  echo "$kernel 版本低于 4.9.17"
  sh upgrade_kernel.sh
  exit 1
fi

if [[ $kernel_v1 -lt 4 ]] && [[ $kernel_v2 -eq 9 ]] &&  [[ $kernel_v3 -lt 17 ]];then
  echo "$kernel 版本低于 4.9.17"
  sh upgrade_kernel.sh
  exit 1
fi

while [ $# -gt 0 ]
do
    key="$1"
    case $key in
        --podCidr)
            export POD_CIDR=$2
            shift
        ;;
    esac
    shift
done


#export CILIUM_VERSION=$(curl -s https://ghproxy.com/https://raw.githubusercontent.com/cilium/cilium/master/stable.txt)
#export HUBBLE_VERSION=$(curl -s https://ghproxy.com/https://raw.githubusercontent.com/cilium/hubble/master/stable.txt)


downloadCilium(){
  curl -L --remote-name-all https://ghproxy.com/https://github.com/cilium/cilium-cli/releases/latest/download/cilium-linux-amd64.tar.gz{,.sha256sum}
  sha256sum --check cilium-linux-amd64.tar.gz.sha256sum
  sudo tar xzvfC cilium-linux-amd64.tar.gz /usr/local/bin
  rm cilium-linux-amd64.tar.gz{,.sha256sum}
  cilium install --config cluster-pool-ipv4-cidr=$POD_CIDR
}

installCilium(){
  helm repo add cilium https://helm.cilium.io/
  helm repo update
  helm install cilium cilium/cilium --version 1.13.2 \
  --namespace kube-system
}

setBgpConfigMap(){
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: bgp-config
  namespace: kube-system
data:
  config.yaml: |
    peers:
      - peer-address: 192.168.0.2
        peer-asn: 64512
        my-asn: 64512
EOF
}

main(){
  installCilium
}

main