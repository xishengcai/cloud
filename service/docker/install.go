package docker

import (
	"fmt"
	"time"

	ssh2 "golang.org/x/crypto/ssh"
	"golang.org/x/net/context"

	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/pkg/sshhelper"
)

const (
	InstallDockerScript    = "/root/install_docker.sh"
	InstallDockerScriptTpl = "./template/install_docker.sh"
)

func InstallDocker(ctx context.Context, client *ssh2.Client) (err error) {
	if err := sshhelper.ScpFile(InstallDockerScriptTpl, InstallDockerScript, client); err != nil {
		return err
	}
	var b []byte
	go func() {
		b, err = sshhelper.ExecCmd(client, "sh /root/install_docker.sh")
	}()
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("安装docker 超时，160秒")
		default:
			if err != nil || len(b) > 0 {
				klog.Infof("install docker resp: %s", string(b))
				klog.Infof("nodes: %s, install docker success", client.RemoteAddr())
				return
			}
			time.Sleep(time.Second * 10)
			klog.Infof("nodes: %s, installing docker ....", client.RemoteAddr())
		}
	}
}
