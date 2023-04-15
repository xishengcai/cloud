# release v1
- install k8s single node

# release v2
- install high available masters
- install slaves

## 2021.5.27
- 重构集群安装代码

## 2022.3.25
- 支持 集群升级（1.21，非外网不可用，因为coredns多了一层父目录）
```
k8s.gcr.io/kube-apiserver:v1.21.11  						4d217480042e
k8s.gcr.io/kube-controller-manager:v1.21.11   	09708983cc37
k8s.gcr.io/kube-scheduler:v1.21.11   						62ad3129eca8
k8s.gcr.io/kube-proxy:v1.21.11     							38ddd85fe90e
k8s.gcr.io/pause:3.4.1													0f8457a4c2ec	
k8s.gcr.io/etcd:3.4.13-0												0369cf4303ff
k8s.gcr.io/coredns/coredns:v1.8.0               296a6d5035e2
```

## 2023.3.15
- remove redis