package images

import (
	"fmt"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/ossutil"
	"github.com/xishengcai/cloud/pkg/setting"
)

type target interface {
	push() error
	url(tarName string) string
}

type OssTarget struct {
	OSS ossutil.AliOssHelper
	Dir string
}

func newOssTarget(dir string) OssTarget {
	return OssTarget{
		OSS: ossutil.NewAliCloudOSS(),
		Dir: dir,
	}
}

func (o OssTarget) push() error {
	return nil
}
func (o OssTarget) url(tarName string) string {
	return fmt.Sprintf("https://%s.%s/%s/%s",
		setting.Config.AliCloud.OSS.Bucket,
		setting.Config.AliCloud.OSS.Endpoint,
		o.Dir,
		tarName)
}

type Registry struct {
	Address string `json:"address"`
}

func (r Registry) push() error {
	return nil
}

func (r Registry) url(tarName string) string {
	return ""
}

type RemoteHost struct {
	models.Host
}

func (r RemoteHost) push() error {
	return nil
}

func (r RemoteHost) url(tarName string) string {
	return ""
}

type Local struct {
	Path string `json:"path" default:"/data/images"`
}

func (l Local) push() error {
	return nil
}

func (l Local) url(tarName string) string {
	return ""
}
func NewTarget(p *Pull) target {
	switch p.Type {
	case PushRemoteHost:
		return p.RemoteHost
	case PushRegistry:
		return p.Registry
	case PushOSS:
		return newOssTarget("/abc")
	default:
		return Local{}
	}
}

type PushType string

const (
	PushRegistry   PushType = "registry"
	PushRemoteHost PushType = "remote_host"
	PushOSS        PushType = "oss"
	PushLocal      PushType = "local"
)
