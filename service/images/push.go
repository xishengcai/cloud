package images

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/pkg/ossutil"
)

type Pull struct {
	Source         Repo                 `json:"source"`
	OSS            ossutil.AliOssHelper `json:"-"`
	*client.Client `json:"-"`
}

func (p *Pull) Validate() error {
	return nil
}

func (p *Pull) Run() app.ResultRaw {
	if p.Source.Addr == "" {
		p.Source.Addr = "docker.io"
	}
	if p.Source.Org == "" {
		p.Source.Org = "library"
	}
	go func() {
		err := p.download()
		if err != nil {
			klog.Error(err)
		}
	}()
	return app.NewServiceResult(nil, nil)
}

func (p *Pull) download() error {
	for _, image := range p.Source.Images {
		fullName := p.Source.Addr + "/" + p.Source.Org + "/" + image.Name + ":" + image.Version
		klog.Info("image fullName: ", fullName)
		reader, err := p.ImagePull(
			context.Background(),
			fullName,
			types.ImagePullOptions{})
		if err != nil {
			klog.Error(err)
			return err
		}
		io.Copy(os.Stdout, reader)
		klog.Infof("image download success")
		shortName := fmt.Sprintf("%s-%s.tar", image.Name, image.Version)
		path, err := p.saveImage(fullName, "/opt/image_ftp")
		os.Remove(path)
		if err != nil {
			os.Remove(path)
			return err
		}
		// 上传Byte数组
		err = p.OSS.Bucket.PutObjectFromFile("idp/"+shortName, path, nil)
		if err != nil {
			os.Remove(path)
			return err
		}
		os.Remove(path)
		klog.Info("image upload to oss success")
	}
	return nil
}

func (p *Pull) saveImage(imageName, dir string) (string, error) {
	path := fmt.Sprintf("%s/%s.tar", dir, uuid.NewUUID())
	w, err := os.Create(path)
	if err != nil {
		return path, err
	}
	r, err := p.ImageSave(context.Background(), []string{imageName})
	if err != nil {
		return path, err
	}
	n, err := io.Copy(w, r)
	if err != nil {
		return path, err
	}
	klog.Infof("image save success, size: %d", n/1024/1024)
	return path, nil
}
