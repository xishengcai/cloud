package kubernetes

import (
	"fmt"

	ssh2 "golang.org/x/crypto/ssh"
	"golang.org/x/net/context"

	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/service/docker"
)

// Version cluster version, used for shell template parser
type Version struct {
	Version string `form:"version" default:"1.22.15"`
}

// JoinNodes batch join node to  k8s
type JoinNodes struct {
	WorkNodes             []models.Host `json:"workNodes"`
	ControllerNodes       []models.Host `json:"controllerNodes"`
	Master                models.Host   `json:"master"`
	Version               string        `json:"-"`
	JoinWorkNodeCommand   string        `json:"-"`
	JoinControllerCommand string        `json:"-"`
}

// Run install slave by master config [hostIP, port root, password]
func (i *JoinNodes) Run() app.ResultRaw {
	jobChan <- i
	klog.Infof("enqueue job: %v", i)
	return app.NewServiceResult(nil, nil)

}

func (i *JoinNodes) startJob() {
	err := i.setJoinCommand()
	if err != nil {
		klog.Error(err)
		return
	}
	i.join()
}

func (i *JoinNodes) Validate() error {
	return nil
}

func (i *JoinNodes) setJoinCommand() error {
	// 如果是和master 一起安装，无需再次登陆获取
	if i.JoinWorkNodeCommand != "" {
		return nil
	}
	client, err := i.Master.GetSSHClient()
	if err != nil {
		return err
	}
	i.JoinWorkNodeCommand, err = getJoinWorkNodeCommand(client)
	if err != nil {
		return err
	}
	i.JoinControllerCommand, err = getJoinControllerNodeCommand(client, i.JoinWorkNodeCommand)
	if err != nil {
		return err
	}
	i.Version, err = getClusterVersion(client)
	return err
}

func (i *JoinNodes) join() {
	for _, node := range i.WorkNodes {
		go func(host models.Host) {
			client, err := host.GetSSHClient()
			if err != nil {
				klog.Error(err)
				return
			}
			ctx, cancel := context.WithTimeout(context.TODO(), installDockerTimeOut)
			defer cancel()
			err = docker.InstallDocker(ctx, client)
			if err != nil {
				klog.Error(err)
				return
			}
			err = joinNode(client, i.Version, i.JoinWorkNodeCommand)
			if err != nil {
				klog.Error(err)
				return
			}
			klog.Infof("node: %s join kubernetes success", host.IP)
		}(node)
	}

	// wait until
	for _, node := range i.ControllerNodes {
		go func(host models.Host) {
			client, err := host.GetSSHClient()
			if err != nil {
				klog.Error(err)
				return
			}
			err = joinNode(client, i.Version, i.JoinControllerCommand)
			if err != nil {
				klog.Error(err)
				return
			}
			klog.Infof("node: %s join kubernetes success", host.IP)
		}(node)
	}
}

func joinNode(client *ssh2.Client, version, joinCmd string) (err error) {
	if err := scpData(client, Version{Version: version}, []string{installKubeletTpl}); err != nil {
		return err
	}

	commands := []string{
		fmt.Sprintf(`sh %s`, targetFile(installKubeletTpl)),
		joinCmd,
	}
	if err := executeCmd(client, commands); err != nil {
		return err
	}
	return nil
}
