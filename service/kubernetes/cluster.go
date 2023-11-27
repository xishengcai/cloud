package kubernetes

import (
	"fmt"

	ssh2 "golang.org/x/crypto/ssh"
	"golang.org/x/net/context"

	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/pkg/common"
	"github.com/xishengcai/cloud/service/docker"
)

const (
	installKubeletTpl         = "./template/install_kubeadm.sh"
	installK8sMasterScriptTpl = "./template/install_k8s_master.sh"
	upgradeKubelet            = "./template/upgrade_nodes.sh"
	ciliumLinuxTpl            = "./template/cilium_linux.sh"
	upgradeKernelShell        = "./template/upgrade_kernel.sh"
)

// Cluster implement install k8s master and slave
// ssh to nodes, run shell script
type Cluster struct {
	*models.Cluster
	JoinWorkNodeCommand   string        `json:"-"`
	JoinControllerCommand string        `json:"-"`
	Skip                  map[step]bool `json:"skip"`
}

func (i *Cluster) Validate() error {
	return nil
}

// Run Install export to API interface
func (i *Cluster) Run() app.ResultRaw {
	i.ID = common.GetUUID()
	_, err := mongoCollection.InsertOne(context.TODO(), i)
	if err != nil {
		return app.NewServiceResult(nil, err)
	}
	jobChan <- i
	klog.Infof("enqueue job: %v", i)
	return app.NewServiceResult(nil, nil)
}

// InstallMaster install k8s master
func (i *Cluster) install() error {
	client, err := i.Master[0].GetSSHClient()
	if err != nil {
		klog.Error(err)
		return err
	}

	err = docker.InstallDocker(client)
	if err != nil {
		klog.Errorf("install docker failed: %v", err)
		return err
	}

	err = i.InstallMaster(client)
	if err != nil {
		klog.Errorf("install master failed: %v", err)
		return err
	}
	err = i.setCommand(client)
	return err
}

func (i *Cluster) NewJoinNodes() JoinNodes {
	jn := JoinNodes{
		WorkNodes: i.WorkNodes,
		Master:    i.Master[0],
		Version:   i.Version,
	}
	for _, x := range i.Master {
		jn.ControllerNodes = append(jn.ControllerNodes, x)
	}
	return jn
}

func (i *Cluster) startJob() {
	err := i.install()
	if err != nil {
		klog.Errorf("install master failed: %v", err)
		return
	}
	if len(i.WorkNodes) > 0 {
		nodes := i.NewJoinNodes()
		nodes.startJob()
	}
}

// InstallMaster kube init by kubeadm_config, or join k8s as master role
func (i *Cluster) InstallMaster(client *ssh2.Client) (err error) {
	if err := scpData(client, i, []string{
		installKubeletTpl,
		installK8sMasterScriptTpl,
		ciliumLinuxTpl,
		upgradeKernelShell}); err != nil {
		return err
	}

	commands := []string{
		fmt.Sprintf(`sh %s`, targetFile(installKubeletTpl)),
		fmt.Sprintf(`sh %s`, targetFile(installK8sMasterScriptTpl)),
		fmt.Sprintf(`cat %s`, "/root/.kube/config"),
	}
	if err := executeCmd(client, commands); err != nil {
		return err
	}
	return
}

func (i *Cluster) setCommand(client *ssh2.Client) error {
	var err error
	if len(i.Master) > 1 || len(i.WorkNodes) > 0 {
		return nil
	}
	i.JoinWorkNodeCommand, err = getJoinWorkNodeCommand(client)
	if err != nil {
		return err
	}
	i.JoinControllerCommand, err = getJoinControllerNodeCommand(client, i.JoinWorkNodeCommand)
	return err
}
