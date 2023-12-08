#!/usr/bin/env bash

installPlugin() {
  # https://artifacthub.io/packages/helm/metrics-server/metrics-server
  helm upgrade --install metrics-server gt/metrics-server --version 3.11.0 \
  --set --image.repository=xisheng/metrics-server \
  --set image.tag=v0.6.4 -n kube-system

  #
  helm upgrade --install dashboard gt/kubernetes-dashboard --version 6.0.8 -n kube-system
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: kubernetes-dashboard
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: admin-user
    namespace: kube-system
EOF
  echo '
  kubectl get secret admin-user-token-p7dwp -n kube-system -o jsonpath={".data.token"} | base64 -d
  '
}

main(){
  installPlugin
}
main