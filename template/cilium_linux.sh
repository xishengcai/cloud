#!/bin/bash

set -e

cilimumTar=https://github.com/cilium/cilium-cli/releases/latest/download/cilium-linux-amd64.tar.gz

checkNet(){
  # -c: 表示次数，1 为1次
  # -w: 表示deadline, time out的时间，单位为秒，100为100秒。
  ping -c 1 -w 3 www.google.com
  if [[ $? != 0 ]];then
    echo "in inter"
    cilimumTar=https://soft-package-xisheng.oss-cn-hangzhou.aliyuncs.com/k8s/cilium/v0.8.4/cilium-linux-amd64.tar.gz
  fi
}

download(){
  curl -LO $cilimumTar
  sudo tar xzvfC cilium-linux-amd64.tar.gz /usr/local/bin
  rm -rf cilium-linux-amd64.tar.gz
}

installCilium(){
  cilium install
}

installCheck(){
  kubectl apply -f https://raw.githubusercontent.com/cilium/cilium/master/examples/kubernetes/connectivity-check/connectivity-check.yaml
}


enableHubble() {
  # Enabling Hubble requires the TCP port 4245 to be open on all nodes running Cilium. T
  # his is required for Relay to operate correctly.
  cilium hubble enable
}

downloadHubbleClient(){
  export HUBBLE_VERSION=$(curl -s https://raw.githubusercontent.com/cilium/hubble/master/stable.txt)
  curl -LO "https://github.com/cilium/hubble/releases/download/$HUBBLE_VERSION/hubble-linux-amd64.tar.gz"
  curl -LO "https://github.com/cilium/hubble/releases/download/$HUBBLE_VERSION/hubble-linux-amd64.tar.gz.sha256sum"
  sha256sum --check hubble-linux-amd64.tar.gz.sha256sum
  tar zxf hubble-linux-amd64.tar.gz

}

main(){
  checkNet
  download
  installCilium
#  enableHubble
}

main