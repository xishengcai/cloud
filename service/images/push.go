package images

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/pkg/app"
)

const (
	ImageCacheDir = "./image_cache"
)

type Pull struct {
	Source     Repo       `json:"source"`
	Type       PushType   `json:"type" default:"local"`
	Registry   Registry   `json:"registry"`
	RemoteHost RemoteHost `json:"remoteHost"`
	Local      Local      `json:"local"`
}

func (p *Pull) Validate() error {
	return nil
}

func (p *Pull) Run() app.ResultRaw {
	pullQueue <- p
	return app.NewServiceResult(nil, nil)
}

func (p *Pull) download() (ImageInfo, error) {
	client, err := p.Source.getClient()
	if err != nil {
		return ImageInfo{}, err
	}
	for _, image := range p.Source.Images {
		imageFullName := p.Source.Addr + "/" + p.Source.Org + "/" + image.Name + ":" + image.Version
		klog.Info("image full name: ", imageFullName)
		imageInfo := ImageInfo{
			Version:  image.Version,
			Org:      p.Source.Org,
			Host:     p.Source.Addr,
			Status:   downloading,
			FullName: imageFullName,
			Updated:  time.Now(),
			Name:     image.Name,
		}
		cache.set(image.Name, imageInfo)
		reader, err := client.ImagePull(
			context.Background(),
			imageFullName,
			types.ImagePullOptions{})
		if err != nil {
			klog.Error(err)
			return imageInfo, err
		}
		io.Copy(os.Stdout, reader)
		klog.Infof("image download success")
		cache.setStatus(imageInfo, saving)

		tarPackage := fmt.Sprintf("%s-%s.tar", image.Name, image.Version)
		imageCache := NewImageCache(imageFullName, tarPackage, ImageCacheDir)
		defer os.Remove(imageCache.getTemplatePath())
		if err = p.saveImage(imageCache); err != nil {
			return imageInfo, err
		}

		// 上传Byte数组
		target := NewTarget(p)
		err = target.push(imageCache)
		if err != nil {
			return imageInfo, err
		}
		cache.setStatus(imageInfo, success)
		cache.setURL(imageInfo, target.url(tarPackage))
		saveToLocal()
	}
	return ImageInfo{}, nil
}

func (p *Pull) saveImage(imageCache ImageCache) error {
	path := imageCache.getTemplatePath()
	w, err := os.Create(path)
	if err != nil {
		return err
	}

	client, err := p.Source.getClient()
	if err != nil {
		return err
	}
	r, err := client.ImageSave(context.Background(), []string{imageCache.FullName})
	if err != nil {
		return err
	}
	n, err := io.Copy(w, r)
	if err != nil {
		return err
	}
	klog.Infof("image save success, size: %d", n/1024/1024)
	return nil
}

type ImageCache struct {
	FullName    string
	TarPackage  string
	TemplateDir string
}

func NewImageCache(fullName, targetPackage, templateDir string) ImageCache {
	return ImageCache{
		FullName:    fullName,
		TarPackage:  targetPackage,
		TemplateDir: templateDir,
	}
}

func (i ImageCache) getTemplatePath() string {
	return fmt.Sprintf("%s/%s", ImageCacheDir, i.TemplateDir)
}
