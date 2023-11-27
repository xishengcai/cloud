package kubernetes

import (
	"time"

	"github.com/pkg/errors"

	"github.com/xishengcai/cloud/models"
)

var (
	errHostPasswordIsNull   = errors.New("not found master host password")
	errMasterIPIsNull       = errors.New("when master enable, master ip can not be null")
	errMasterPasswordIsNull = errors.New("when master enable, master password can not be null")
	errNodesIsNull          = errors.New("nodes is null")
	joinCache               *JoinCache
)

type JoinCache struct {
	JoinNodeCommand string
	Version         string
	RegisterTime    time.Time
}

func NewJoinCache() (*JoinCache, error) {
	if joinCache == nil {
		return newJoinCache()
	}
	if joinCache.RegisterTime.Second()+3600 < time.Now().Second() {
		return newJoinCache()
	}
	return joinCache, nil
}

func newJoinCache() (*JoinCache, error) {
	var host models.Host
	err := models.GetDB("host").Where("is_master = ?", true).Find(&host).Error
	if err != nil {
		return nil, err
	}

	if host.Password == "" {
		return nil, errHostPasswordIsNull
	}

	client, err := host.GetSSHClient()
	if err != nil {
		return nil, err
	}

	joinNodeCommand, err := getJoinNodeCommand(client)
	if err != nil {
		return nil, err
	}

	version, err := getClusterVersion(client)
	if err != nil {
		return nil, err
	}
	joinCache = &JoinCache{
		JoinNodeCommand: joinNodeCommand,
		Version:         version,
		RegisterTime:    time.Now(),
	}
	return joinCache, err
}
