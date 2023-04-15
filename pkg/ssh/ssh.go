package ssh

import (
	"fmt"
	"net"
	"time"

	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/models"

	"github.com/pkg/errors"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
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
	if err != nil {
		return err
	}
	klog.Info("copy byteStream to remote server finished!")
	return nil
}

// GetClient 通过ssh.ClientConfig创建一个ssh连接
func GetClient(host models.Host) (*ssh.Client, error) {
	klog.InfoS("host message", "host", host, "user", host.User, "password", host.Password, "port", host.Port)
	addr := fmt.Sprintf("%s:%d", host.IP, host.Port)
	sshConfig := getConfigByPassword(host.User, host.Password, time.Second*5)
	client, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, errors.Wrap(err, "dial to host")
	} else if client == nil {
		return nil, errors.Wrap(err, "get ssh client")
	}
	return client, nil
}
