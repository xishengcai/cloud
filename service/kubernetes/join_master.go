package kubernetes

import (
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/pkg/e"
	"github.com/xishengcai/cloud/pkg/ssh"

	"k8s.io/klog/v2"
)

// JoinMaster batch join master to  k8s
type JoinMaster struct {
	InstallSlave
}

// Run install slave by master config [hostIP, port root, password]
func (i JoinMaster) Run() app.ResultRaw {
	err := i.install()
	return app.NewServiceResultWitRawData(i, err)

}

func (i JoinMaster) startJob() {
	err := i.install()
	if err != nil {
		klog.Error(err)
	}
}

func (i JoinMaster) install() error {
	if err := i.getVersion(i.Master); err != nil {
		return err
	}

	if err := i.setJoinCommand(); err != nil {
		return err
	}

	jobChan <- i
	return nil
}

func (i JoinMaster) setJoinCommand() error {
	client, err := ssh.GetClient(i.Master)
	if err != nil {
		klog.Error(err)
		return err
	}
	joinCommand, err := getJoinMasterCommand(client)
	if err != nil {
		return err
	}
	i.JoinSlaveCommand = joinCommand
	return nil
}

func (i JoinMaster) joinNodes() (err error) {
	var errorList []error
	for _, item := range i.Nodes {
		err := joinNode(item, i.Version, i.JoinSlaveCommand, i.DryRun)
		if err != nil {
			errorList = append(errorList, err)
		}
	}
	// return fmt.Error(errors.NewAggregate(errorList).Error())
	return e.MergeError(errorList)
}
