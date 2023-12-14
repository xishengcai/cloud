# function

- [x] install high available k8s masters (需要自己实现vip的负载均衡)
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
    - [x] 1.20.xxx
    - [] 1.21.xxx
    - [x] 1.22.xxx
    - [x] 1.23.xxx
  
- [x] upgrade k8s
    - [x] 1.17 ---> 1.18.x
    - [x] 1.18 ---> 1.19.x
    - [x] 1.19 ---> 1.20.x
    - [x] 1.20.x --> 1.21.0 升级 (特殊)
    - [x] 1.21.x --> 1.22.x 升级 
    - [x] 1.22.x --> 1.23.x 升级 
  
## local start
```bash
docker run -d --restart=always -p 27017:27017 --name mongodb -v /var/lib/mongo:/data/mongodb \
-e MONGO_INITDB_ROOT_PASSWORD=123456 -e MONGO_INITDB_DATABASE=admin -e MONGO_INITDB_ROOT_USERNAME=root \
-d mongo:latest

go run main.go

cd dashboard && npm run dev
```

## response code
```
- 200, 表示操作成功
- 400, 参数不符合
- 201, 任务提交成功
- 0, 业务处理失败

```
![img.png](img.png)