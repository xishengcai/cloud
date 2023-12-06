package sshhelper

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"text/template"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"
)

// getConfigByPassword 通过用户名和密码生成一个配置文件
func getConfigByPassword(user string, password string, timeout time.Duration) *ssh.ClientConfig {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: timeout,
	}
	return sshConfig
}

// ExecCmd 通过*ssh.Client 执行命令
func ExecCmd(client *ssh.Client, cmd string) ([]byte, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = session.Close()
	}()
	out, err := session.CombinedOutput(cmd)
	return out, err
}

// CopyByteToRemote 复制字节数组到远程服务器上
func CopyByteToRemote(client *ssh.Client, byteStream []byte, remoteFilePath string) error {
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		return err
	}
	defer func() {
		_ = sftpClient.Close()
	}()

	// create directory
	dstFile, err := sftpClient.Create(remoteFilePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = dstFile.Close()
	}()
	_, err = dstFile.Write(byteStream)
	return err
}

func GetClient(ip, user, password string, port int) (*ssh.Client, error) {
	klog.InfoS("host message", "ip", ip, "user", user, "password", password, "port", port)
	addr := fmt.Sprintf("%s:%d", ip, port)
	sshConfig := getConfigByPassword(user, password, time.Second*5)
	client, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		klog.Error(err)
	}
	return client, err
}

// ScpData use go text template, data is a struct object
// temp is shell script, struct element use {{.Element}}
func ScpData(client *ssh.Client, data interface{}, pathMap map[string]string) error {
	for k, v := range pathMap {
		klog.Infof("srcPath: %s, destPath: %s", k, v)
		scriptBytes, err := parserTemplate(k, data)
		if err != nil {
			return err
		}
		if err := CopyByteToRemote(client, scriptBytes, v); err != nil {
			return err
		}
	}
	return nil
}

func TargetFile(tmp string) string {
	t := strings.Split(tmp, "/")
	return "/root/" + t[len(t)-1]
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
