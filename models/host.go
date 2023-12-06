package models

import (
	ssh2 "golang.org/x/crypto/ssh"

	"github.com/xishengcai/cloud/pkg/sshhelper"
)

type Host struct {
	IP       string `json:"ip" bson:"ip"`
	User     string `json:"user" default:"root" bson:"user"`
	Port     int    `json:"port" bson:"port" default:"22"`
	Password string `json:"password" bson:"password"`
}

func (h *Host) GetSSHClient() (*ssh2.Client, error) {
	return sshhelper.GetClient(h.IP, "root", h.Password, h.Port)
}
