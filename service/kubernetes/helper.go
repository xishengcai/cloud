package kubernetes

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	ssh2 "golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/pkg/ssh"
)

func parserTemplate(scriptTpl string, data interface{}) ([]byte, error) {
	t1, err := template.ParseFiles(scriptTpl)
	if err != nil {
		klog.Errorf("%s template parser failed, %v", scriptTpl, err)
		return nil, err
	}
	buff1 := new(bytes.Buffer)

	// 结构体数据映射到模版中
	err = t1.Execute(buff1, data)
	if err != nil {
		klog.Errorf("execute template failed, %v", err)
		return nil, err
	}
	return buff1.Bytes(), nil
}

// scpData use go text template, data is a struct object
// temp is shell script, struct element use {{.Element}}
func scpData(client *ssh2.Client, data interface{}, temp []string) error {
	for _, t := range temp {
		scriptBytes, err := parserTemplate(t, data)
		if err != nil {
			return err
		}
		if err := ssh.CopyByteToRemote(client, scriptBytes, targetFile(t)); err != nil {
			return err
		}
	}
	return nil
}

func executeCmd(client *ssh2.Client, commands []string) error {
	for _, cmd := range commands {
		klog.Infof("exec cmd %s", cmd)
		b, err := ssh.ExecCmd(client, cmd)
		klog.Infof("resp:  %s", string(b))
		if err != nil {
			klog.Errorf("ExecCmd failed, %v", err)
			return err
		}
		klog.Infof("exec cmd: %s success", cmd)
	}
	return nil
}

func targetFile(tmp string) string {
	t := strings.Split(tmp, "/")
	return "/root/" + t[len(t)-1]
}

// getClusterVersion ssh to kubernetes master, get version
func getClusterVersion(client *ssh2.Client) (string, error) {
	b, err := ssh.ExecCmd(client, "kubectl version --short |grep Server")
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
	result, err := ssh.ExecCmd(client, "kubeadm init phase upload-certs --upload-certs | awk 'END {print}'")
	if err != nil {
		return "", err
	}
	certificateKey := handCommandResult(result)

	return joinWorkNodeCommand + " --control-plane --certificate-key  " + certificateKey, nil

}

func getJoinWorkNodeCommand(client *ssh2.Client) (string, error) {
	b, err := ssh.ExecCmd(client, "kubeadm token create --print-join-command")
	if err != nil {
		return "", err
	}
	cmdStr := string(b)
	cmd := cmdStr[strings.Index(cmdStr, "kubeadm join"):]
	klog.Infof("node join command: %s", cmd)
	return cmd, nil
}

func getJoinNodeCommand(client *ssh2.Client) (string, error) {
	b, err := ssh.ExecCmd(client, "kubeadm token create --print-join-command")
	if err != nil {
		return "", err
	}
	cmdStr := string(b)
	cmd := cmdStr[strings.Index(cmdStr, "kubeadm join"):]
	klog.Infof("node join command: %s", cmd)
	return cmd, nil
}
