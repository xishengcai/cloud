set -e
yum remove kubelet -y
yum install -y kubelet-{{.Version}}  kubeadm-{{.Version}}  kubectl-{{.Version}}
systemctl daemon-reload
systemctl restart kubelet

if [ -f "/root/.kube/config" ]; then
  kubectl -n kube-system get cm kubeadm-config --template={{.data.ClusterConfiguration}} > upgrade-k8s.yaml
  sed -i 's/^kubernetesVersion:.*/kubernetesVersion: {{.Version}}/g' upgrade-k8s.yaml
  if [ "{{.Version}}" = "1.21.0" ]; then
    wget -O /usr/bin/kubeadm https://lstack-qa.oss-cn-hangzhou.aliyuncs.com/kubeadm-{{.Version}}
  end


  kubeadm upgrade apply --config upgrade-k8s.yaml -y
fi

