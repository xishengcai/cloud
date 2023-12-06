package proxy

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/pkg/ssh"
)

const (
	proxyNgConfigDistPath = "/etc/nginx/conf.d/proxy.conf"
	proxyNgConfigSrc      = "./template/v2ray/proxy.conf"
)

const (
	v2rayVmessTLSWSConfigDistPath = "/etc/v2ray/vmess-tls-ws.json"
	v2rayVmessTLSWSConfigSrc      = "./template/v2ray/vmess-tls-ws.json"
)

const (
	certSrcPath  = "./template/gencert"
	certDistPath = "/root/gencert"
)

type Install struct {
	models.Host
	CommonName   string `json:"commonName"`
	ExternalPort int    `json:"externalPort"` // used for nginx tls， listen to v2ray
	V2rayPort    int    `json:"v2rayPort"`
}

func (i *Install) Validate() error {
	return nil
}

func (i *Install) Run() app.ResultRaw {
	err := i.installV2ray(time.Second * 120)
	return app.NewServiceResult(nil, err)
}

func (i *Install) installV2ray(duration time.Duration) (err error) {
	client, err := i.GetSSHClient()
	if err != nil {
		klog.Error(err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.TODO(), duration)
	defer cancel()
	if err = ssh.ScpFile(proxyNgConfigSrc, proxyNgConfigDistPath, client); err != nil {
		return err
	}

	if err = ssh.ScpFile(v2rayVmessTLSWSConfigSrc, v2rayVmessTLSWSConfigDistPath, client); err != nil {
		return err
	}

	if err = ssh.ScpFile(certSrcPath, certDistPath, client); err != nil {
		return err
	}

	var b []byte
	go func() {
		b, err = ssh.ExecCmd(client, v2rayScript)
	}()
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("安装proxy失败")
		default:
			if err != nil || len(b) > 0 {
				klog.Infof("安装v2ray resp: %s", string(b))
				klog.Infof("节点: %s, 安装v2ray 成功", client.RemoteAddr())
				return
			}
			time.Sleep(time.Second * 10)
			klog.Infof("节点: %s, 正在安装v2ray ....", client.RemoteAddr())
		}
	}
}

const v2rayScript = `
#!/bin/bash
wget v2ray
yum install nginx wget -y
wget https://github.com/xishengcai/cloud/releases/download/v1.0.0/proxy-server
chmod +x /root/proxy-server
set -e
sh /root/gencert --CN test.hello.com --dir=/opt/proxy-cert
/root/proxy-server --config /etc/v2ray/vmess-tls-ws.json &
systemctl start nginx
`
