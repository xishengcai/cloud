package kubernetes

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/pkg/e"
	"github.com/xishengcai/cloud/pkg/ssh"
	"github.com/xishengcai/cloud/service/docker"

	"github.com/pkg/errors"

	ssh2 "golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"
)

const (
	installKubeletTpl         = "./template/install_kubeadm.sh"
	installK8sMasterScriptTpl = "./template/install_k8s_master.sh"
	upgradeKubelet            = "./template/upgrade_nodes.sh"
	calicoYamlTpl             = "./template/calico.yaml"
	ciliumLinuxTpl            = "./template/cilium_linux.sh"
	flannelTpl                = "./template/flannel.yaml"
)

const (
	calico  = "calico"
	cilium  = "cilium"
	flannel = "flannel"
)

var (
	networkPlugin = map[string]string{
		calico:  calicoYamlTpl,
		cilium:  ciliumLinuxTpl,
		flannel: flannelTpl,
	}
)

// InstallKuber implement install k8s master and slave
// ssh to nodes, run shell script
type InstallKuber struct {
	*models.Kubernetes
	DryRun bool `json:"dryRun"`
}

func (i InstallKuber) Validate() error {
	return nil
}

// Run Install export to API interface
func (i InstallKuber) Run() app.ResultRaw {
	jobChan <- i
	klog.Infof("enqueue job: %v", i)
	return app.NewServiceResult(nil, nil)
}

// InstallMaster install k8s master
func (i InstallKuber) install() error {
	client, err := ssh.GetClient(i.PrimaryMaster)
	if err != nil {
		klog.Error(err)
		return err
	}
	err = docker.InstallDocker(i.PrimaryMaster, i.DryRun)
	if err != nil {
		klog.Errorf("install docker failed: %v", err)
		return errors.Wrap(err, "install docker")
	}

	err = i.InstallMaster(i.PrimaryMaster)
	if err != nil {
		klog.Errorf("install master failed: %v", err)
		return errors.Wrap(err, "install master")
	}

	// get joinMaster cmd
	i.JoinMasterCommand, err = getJoinMasterCommand(client)
	if err != nil {
		klog.Errorf("getJoinMasterCommand failed: %v", err)
		return errors.Wrap(err, "getJoinMasterCommand failed")
	}

	klog.Infof("joinMasterCommand: %s", i.JoinMasterCommand)
	var errs []error
	for _, item := range i.BackendMasters {
		err = joinNode(item, i.Version, i.JoinMasterCommand, i.DryRun)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return e.MergeError(errs)
}

func (i InstallKuber) startJob() {
	err := i.install()
	if err != nil {
		klog.Error(err)
	}
}

// InstallMaster kube init by kubeadm_config, or join k8s as master role
func (i InstallKuber) InstallMaster(host models.Host) (err error) {
	client, err := ssh.GetClient(host)
	if err != nil {
		return err
	}
	networkPluginTpl := networkPlugin[i.NetWorkPlug]

	var installNetPluginCmd string
	if i.NetWorkPlug == calico {
		installNetPluginCmd = "kubectl apply -f /root/calico.yaml"
	} else {
		installNetPluginCmd = "sh cilium_linux.sh"
	}
	if err := scpData(client, i, []string{installKubeletTpl, installK8sMasterScriptTpl, networkPluginTpl}); err != nil {
		return err
	}

	if i.DryRun {
		return nil
	}

	commands := []string{
		fmt.Sprintf(`sh %s`, targetFile(installKubeletTpl)),
		fmt.Sprintf(`sh %s`, targetFile(installK8sMasterScriptTpl)),
		fmt.Sprintf(`cat %s`, "/root/.kube/config"),
		installNetPluginCmd,
	}
	if err := executeCmd(client, commands); err != nil {
		return err
	}
	return
}

func getJoinMasterCommand(client *ssh2.Client) (string, error) {
	jointNodeCmd, err := ssh.ExecCmd(client, "kubeadm token create --print-join-command")
	if err != nil {
		return "", err
	}
	result, err := ssh.ExecCmd(client, "kubeadm init phase upload-certs --upload-certs | awk 'END {print}'")
	if err != nil {
		return "", err
	}
	certificateKey := handCommandResult(result)

	return handCommandResult(jointNodeCmd) + " --control-plane --certificate-key  " + certificateKey, nil

}

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
