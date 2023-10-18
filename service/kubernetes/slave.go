package kubernetes

import (
	"fmt"
	"strings"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/pkg/e"
	"github.com/xishengcai/cloud/pkg/ssh"
	"github.com/xishengcai/cloud/service/docker"

	"k8s.io/klog/v2"
)

// InstallSlave batch join slave to  k8s
type InstallSlave struct {
	Nodes            []models.Host `json:"nodes"`
	Master           models.Host   `json:"master"`
	Version          string        `json:"-"`
	JoinSlaveCommand string        `json:"-"`
	DryRun           bool          `json:"dryRun,omitempty"`
}

// Run install slave by master config [hostIP, port root, password]
func (i *InstallSlave) Run() app.ResultRaw {
	jobChan <- i
	klog.Infof("enqueue job: %v", i)
	return app.NewServiceResult(nil, nil)

}

func (i *InstallSlave) startJob() {
	err := i.joinNodes()
	if err != nil {
		klog.Error(err)
	}
}

func (i *InstallSlave) Validate() error {
	return nil
}

func handCommandResult(result []byte) string {
	slice := strings.Split(string(result), "\n")
	var command string
	if len(slice) >= 1 {
		command = slice[len(slice)-2]
	}
	return command
}

func (i *InstallSlave) getVersion(host models.Host) error {
	client, err := ssh.GetClient(host)
	if err != nil {
		return err
	}
	b, err := ssh.ExecCmd(client, "kubectl version --short |grep Server")
	if err != nil {
		return err
	}
	if len(b) == 0 {
		return fmt.Errorf("not find k8s version from master")
	}
	version := strings.Split(string(b), " ")[2]
	version = strings.Trim(version, "\n")
	version = strings.Trim(version, "v")
	klog.Info("find master version is ", version)
	i.Version = version
	return nil
}

func (i *InstallSlave) setJoinCommand() error {
	joinCommand, err := getJoinNodeCommand(i.Master)
	if err != nil {
		return err
	}

	cmdResp := string(joinCommand)
	joinCMD[i.Master.IP] = cmdResp[strings.Index(cmdResp, "kubeadm join"):]
	return nil
}

func (i *InstallSlave) joinNodes() (err error) {
	var errorList []error
	err = i.setJoinCommand()
	if err != nil {
		return err
	}
	// wait until
	for _, node := range i.Nodes {
		go func(host models.Host) {
			err := joinNode(host, i.Version, joinCMD[i.Master.IP], i.DryRun)
			if err != nil {
				errorList = append(errorList, err)
			}
		}(node)
	}
	// return fmt.Error(errors.NewAggregate(errorList).Error())
	return e.MergeError(errorList)
}

func getJoinNodeCommand(host models.Host) ([]byte, error) {
	client, err := ssh.GetClient(host)
	if err != nil {
		return nil, err
	}
	return ssh.ExecCmd(client, "kubeadm token create --print-join-command")
}

func joinNode(host models.Host, version, joinCmd string, dryRun bool) (err error) {
	client, err := ssh.GetClient(host)
	if err != nil {
		return
	}
	err = docker.InstallDocker(host, dryRun)
	if err != nil {
		return err
	}

	if err := scpData(client, models.Version{Version: version}, []string{installKubeletTpl}); err != nil {
		return err
	}

	if dryRun {
		return nil
	}
	klog.Info("join cmd: ", joinCMD)
	commands := []string{
		fmt.Sprintf(`sh %s`, targetFile(installKubeletTpl)),
		joinCmd,
	}
	if err := executeCmd(client, commands); err != nil {
		return err
	}
	klog.Infof("join node:%s success", host.IP)
	return nil
}
