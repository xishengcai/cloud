package proxy

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/pkg/sshhelper"
)

const (
	proxyNgConfigDestPath = "/etc/nginx/proxy-1234.conf"
	proxyNgConfigSrc      = "./template/proxy/proxy.conf"
)

const (
	proxyVmessTLSWSConfigDestPath = "/root/vmess-tls-ws.json"
	proxyVmessTLSWSConfigSrc      = "./template/proxy/vmess-tls-ws.json"
)

const (
	certSrcPath          = "./template/gencert"
	certDestPath         = "/root/gencert"
	installProxyTpl      = "./template/proxy/install_proxy.sh"
	installProxyDestPath = "/root/install_proxy.sh"
)

var (
	proxyFileMap = map[string]string{
		certSrcPath:     certDestPath,
		installProxyTpl: installProxyDestPath,
	}
)

type Install struct {
	models.Host
	CommonName   string `json:"commonName" default:"test.hello.com"`
	ExternalPort int    `json:"externalPort" default:"20001"` // used for nginx tls， listen to proxy
	ProxyPort    int    `json:"proxyPort" default:"20000"`
}

func (i *Install) Validate() error {
	return nil
}

func (i *Install) Run() app.ResultRaw {
	err := i.installProxy(time.Second * 120)
	return app.NewServiceResult(nil, err)
}

func (i *Install) installProxy(duration time.Duration) (err error) {
	client, err := i.GetSSHClient()
	if err != nil {
		klog.Error(err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.TODO(), duration)
	defer cancel()

	if err := sshhelper.ScpData(client, i, proxyFileMap); err != nil {
		return err
	}
	var b []byte
	go func() {
		b, err = sshhelper.ExecCmd(client, "sh /root/install_proxy.sh")
	}()
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("安装proxy失败")
		default:
			if err != nil || len(b) > 0 {
				klog.Infof("安装proxy resp: %s", string(b))
				klog.Infof("节点: %s, 安装proxy 成功", client.RemoteAddr())
				return
			}
			time.Sleep(time.Second * 10)
			klog.Infof("节点: %s, 正在安装proxy ....", client.RemoteAddr())
		}
	}
}
