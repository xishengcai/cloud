package docker

import (
	"fmt"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/ssh"

	"k8s.io/klog/v2"
)

const (
	InstallDockerScript    = "/root/install_docker.sh"
	InstallDockerScriptTpl = "./template/install_docker.sh"
)

func InstallDocker(host models.Host, dryRun bool) (err error) {
	client, err := ssh.GetClient(host)
	if err != nil {
		return fmt.Errorf("nodes: %s, %v", host.IP, err)
	}

	if err := ssh.ScpFile(InstallDockerScriptTpl,
		InstallDockerScript, client); err != nil {
		return fmt.Errorf("nodes: %s, %v", host.IP, err)
	}

	if dryRun {
		return nil
	}

	b, err := ssh.ExecCmd(client, "sh /root/install_docker.sh")
	if err != nil {
		return fmt.Errorf("nodes: %s, %v", host.IP, err)
	}
	klog.Infof("install docker resp: %s", string(b))
	klog.Infof("nodes: %s, install docker success", host.IP)
	return nil
}
