package kubernetes

import (
	"fmt"
	"strings"

	ssh2 "golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/pkg/sshhelper"
)

func executeCmd(client *ssh2.Client, commands []string) error {
	for _, cmd := range commands {
		klog.Infof("exec cmd %s", cmd)
		b, err := sshhelper.ExecCmd(client, cmd)
		klog.Infof("resp:  %s", string(b))
		if err != nil {
			klog.Errorf("ExecCmd failed, %v", err)
			return err
		}
		klog.Infof("exec cmd: %s success", cmd)
	}
	return nil
}

// getClusterVersion ssh to kubernetes master, get version
func getClusterVersion(client *ssh2.Client) (string, error) {
	b, err := sshhelper.ExecCmd(client, "kubectl version --short |grep Server")
	if err != nil {
		return "", err
	}
	if len(b) == 0 {
		return "", fmt.Errorf("not find k8s version from master")
	}
	version := strings.Split(string(b), " ")[2]
	version = strings.Trim(version, "\n")
	version = strings.Trim(version, "v")
	klog.Info("kubernetes version is ", version)
	return version, nil
}

func handCommandResult(result []byte) string {
	slice := strings.Split(string(result), "\n")
	var command string
	if len(slice) >= 1 {
		command = slice[len(slice)-2]
	}
	return command
}

func getJoinControllerNodeCommand(client *ssh2.Client, joinWorkNodeCommand string) (string, error) {
	result, err := sshhelper.ExecCmd(client, "kubeadm init phase upload-certs --upload-certs | awk 'END {print}'")
	if err != nil {
		return "", err
	}
	certificateKey := handCommandResult(result)

	return strings.TrimRight(joinWorkNodeCommand, "\n") + " --control-plane --certificate-key  " + certificateKey, nil

}

func getJoinWorkNodeCommand(client *ssh2.Client) (string, error) {
	b, err := sshhelper.ExecCmd(client, "kubeadm token create --print-join-command")
	if err != nil {
		return "", err
	}
	cmdStr := string(b)
	cmd := cmdStr[strings.Index(cmdStr, "kubeadm join"):]
	klog.Infof("node join command: %s", cmd)
	return cmd, nil
}

func getJoinNodeCommand(client *ssh2.Client) (string, error) {
	b, err := sshhelper.ExecCmd(client, "kubeadm token create --print-join-command")
	if err != nil {
		return "", err
	}
	cmdStr := string(b)
	cmd := cmdStr[strings.Index(cmdStr, "kubeadm join"):]
	klog.Infof("node join command: %s", cmd)
	return cmd, nil
}
