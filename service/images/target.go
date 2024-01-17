package images

import (
	"fmt"
	"log"
	"os/exec"

	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/ossutil"
	"github.com/xishengcai/cloud/pkg/setting"
)

type target interface {
	push(imageCache ImageCache) error
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

func (o OssTarget) push(imageCache ImageCache) error {
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

func (r Registry) push(imageCache ImageCache) error {
	return nil
}

func (r Registry) url(tarName string) string {
	return ""
}

type RemoteHost struct {
	models.Host
}

func (r RemoteHost) push(imageCache ImageCache) error {
	return nil
}

func (r RemoteHost) url(tarName string) string {
	return ""
}

type Local struct {
	SavePath string `json:"path" default:"/data/images"`
}

func (l Local) push(imageCache ImageCache) error {
	cmd := exec.Command("cp", imageCache.getTemplatePath(), l.SavePath+"/"+imageCache.TarPackage)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to cp image tar: %v", err)
	}
	klog.Infof("push image to local path: /data/images/%s")

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
	case PushLocal:
		return p.Local
	default:
		return nil
	}
}

type PushType string

const (
	PushRegistry   PushType = "registry"
	PushRemoteHost PushType = "remote_host"
	PushOSS        PushType = "oss"
	PushLocal      PushType = "local"
)
