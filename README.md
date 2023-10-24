# function

- [x] install high available k8s masters
- [x] batch install slaves
- [] web terminal
- [x] dashboard(vue)
- [] support skip step, if docker already install
- [] support that when task failed, retry from failed step

## swagger url
- http://xxx:xxx/swagger/index.html

## post example
### install kubernetes
```json
{
  "controlPlaneEndpoint": "string",
  "master": [
    {
      "ip": "string",
      "password": "string",
      "port": 22,
      "user": "root"
    }
  ],
  "name": "test",
  "networkPlug": "cilium",
  "podCidr": "10.244.0.0/16",
  "registry": "registry.aliyuncs.com/google_containers",
  "serviceCidr": "10.96.0.0/16",
  "version": "1.22.15",
  "workNodes": [
    {
      "ip": "string",
      "password": "string",
      "port": 22,
      "user": "root"
    }
  ]
}
```

### join nodes
```json
{
  "controllerNodes": [
    {
      "ip": "string",
      "password": "string",
      "port": 22,
      "user": "root"
    }
  ],
  "master": {
    "ip": "string",
    "password": "string",
    "port": 22,
    "user": "root"
  },
  "skip": {
    "additionalProp1": true,
    "additionalProp2": true,
    "additionalProp3": true
  },
  "version": "string",
  "workNodes": [
    {
      "ip": "string",
      "password": "string",
      "port": 22,
      "user": "root"
    }
  ]
}

```

### pull image and push to oss
```json
{
  "source": {
    "addr": "registry.cn-hangzhou.aliyuncs.com",
    "images": [
      {
        "name": "openjdk",
        "version": "latest"
      }
    ],
    "org": "launcher",
    "password": "",
    "user": ""
  }
}

```

## feature:
- [x] support centos version
    - [x]  CentOS Linux 7 (Core)   3.10.0-957.el7.x86_64
    - [x]  Anolis OS 8.2           4.18.0-193.60.2.an8_2.x86_64
    
- [x] support k8s version
    - [x] 1.17.xxx
    - [x] 1.18.xxx
    - [x] 1.19.xxx
  
- [x] upgrade k8s
    - [x] 1.17 ---> 1.18.x
    - [x] 1.18 ---> 1.19.x
    - [x] 1.19 ---> 1.20.x
    - [x] 1.20.x --> 1.21.0 升级 (特殊)
    - [x] 1.21.x --> 1.22.x 升级 
  
## local start
- go run main.go
- cd front && ng serve --open
- docker run 
```bash

docker run -d -p 8080:8080 -v /user/bin/docker:/user/bin/docker \
-v /etc/docker:/etc/docker \
-v /data/docker:/data/docker \
-v /var/run/docker.sock:/var/run/docker.sock \
-v /opt/image_ftp:/opt/image_ftp \
registry.cn-hangzhou.aliyuncs.com/xisheng/cloud:release-V1
```
![img.png](img.png)