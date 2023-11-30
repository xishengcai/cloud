package proxy

import (
	"golang.org/x/crypto/ssh"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
)

type Install struct {
	models.Host
}

func (i *Install) Validate() error {
	return nil
}

func (i *Install) Run() app.ResultRaw {
	return nil
}

func installV2ray(client *ssh.Client) {

}

const vmessTemplate = ``
const CountResourceShellScript = `
#!/bin/bash
wget v2ray

`
