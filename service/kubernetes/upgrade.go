package kubernetes

import (
	"fmt"
	"strings"

	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/pkg/sshhelper"
)

type Upgrade struct {
	Nodes   []models.Host `json:"nodes"`
	Version string        `json:"version"`
	DryRun  bool          `json:"dryRun"`
}

func (u Upgrade) Validate() error {
	return nil
}

func (u Upgrade) Run() app.ResultRaw {
	jobChan <- u
	return app.NewServiceResult(nil, nil)
}

func (u Upgrade) startJob() {
	err := u.upgradeNodes()
	if err != nil {
		klog.Error(err)
	}
}

func (u Upgrade) upgradeNodes() error {
	for _, h := range u.Nodes {
		go func(host models.Host) {
			err := upgradeNode(host, u.Version, u.DryRun)
			if err != nil {
				klog.Error("upgrade slave err: %v", err)
			}
		}(h)
	}
	return nil
}

func upgradeNode(host models.Host, version string, dryRun bool) error {
	client, err := host.GetSSHClient()
	if err != nil {
		return err
	}

	// TODO: 需要优化
	// 因为kubectl 命令使用了--template={{.data.ClusterConfiguration}}， 在使用go template 解析
	// 会发生解析失败，所以这里使用replace
	script := `
set -e
yum remove kubelet -y

yum install -y kubelet-{{.Version}}  kubeadm-{{.Version}}  kubectl-{{.Version}} wget

if [ -f "/root/.kube/config" ]; then
  kubectl -n kube-system get cm kubeadm-config --template={{.data.ClusterConfiguration}} > upgrade-k8s.yaml
  sed -i 's/^kubernetesVersion:.*/kubernetesVersion: {{.Version}}/g' upgrade-k8s.yaml
  if [ "{{.Version}}" = "1.21.0" ]; then
    wget -O /usr/bin/kubeadm https://lstack-qa.oss-cn-hangzhou.aliyuncs.com/kubeadm-{{.Version}}
  fi
  kubeadm upgrade apply --config upgrade-k8s.yaml -y
fi

systemctl daemon-reload
systemctl restart kubelet
`
	script = strings.Replace(script, "{{.Version}}", version, -1)
	if err := sshhelper.CopyByteToRemote(client, []byte(script), sshhelper.TargetFile(upgradeKubelet)); err != nil {
		return err
	}

	if dryRun {
		return nil
	}
	commands := []string{
		fmt.Sprintf(`sh %s`, sshhelper.TargetFile(upgradeKubelet)),
	}
	if err := executeCmd(client, commands); err != nil {
		return err
	}
	return nil
}
