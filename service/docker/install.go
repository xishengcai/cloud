package docker

import (
	"fmt"

	ssh2 "golang.org/x/crypto/ssh"

	"github.com/xishengcai/cloud/pkg/ssh"

	"k8s.io/klog/v2"
)

const (
	InstallDockerScript    = "/root/install_docker.sh"
	InstallDockerScriptTpl = "./template/install_docker.sh"
)

func InstallDocker(client *ssh2.Client) (err error) {
	if err := ssh.ScpFile(InstallDockerScriptTpl, InstallDockerScript, client); err != nil {
		return fmt.Errorf("nodes: %s, %v", client.RemoteAddr(), err)
	}
	klog.Infof("copy %s to remote server finished!", InstallDockerScriptTpl)

	b, err := ssh.ExecCmd(client, "sh /root/install_docker.sh")
	if err != nil {
		return fmt.Errorf("nodes: %s, %v", client.RemoteAddr(), err)
	}
	klog.Infof("install docker resp: %s", string(b))
	klog.Infof("nodes: %s, install docker success", client.RemoteAddr())
	return nil
}
