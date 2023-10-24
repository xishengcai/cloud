package kubernetes

import (
	ssh "golang.org/x/crypto/ssh"

	"github.com/xishengcai/cloud/service/docker"
)

type step string

const (
	stepSetCommand     step = "SET_COMMAND"
	stepGetVersion     step = "GET_VERSION"
	stepGetJoinCommand step = "GET_JOIN_COMMAND"
	stepInstallDocker  step = "INSTALL_DOCKER"
	stepInstallKubeadm step = "INSTALL_KUBEADM"
)

type TaskInstallDocker struct {
	client *ssh.Client
}
type TaskInstallKubeadm struct {
	client *ssh.Client
}
type TaskInstallNetPlug struct {
	client *ssh.Client
}

func (t TaskInstallDocker) Run() error {
	return docker.InstallDocker(t.client)
}
