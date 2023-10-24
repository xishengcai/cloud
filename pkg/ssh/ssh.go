package ssh

import (
	"fmt"
	"net"
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
