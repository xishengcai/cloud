package ossutil

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/xishengcai/cloud/pkg/setting"
)

type AliOssHelper struct {
	*oss.Bucket
}

func NewAliCloudOSS() AliOssHelper {
	/*
	   oss 的相关配置信息
	*/
	endpoint := setting.Config.AliCloud.OSS.Endpoint
	accessKeyId := setting.Config.AliCloud.OSS.AK
	accessKeySecret := setting.Config.AliCloud.OSS.SK

	//创建OSSClient实例
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}

	var bucket *oss.Bucket
	bucket, err = client.Bucket(setting.Config.AliCloud.OSS.Bucket)

	return AliOssHelper{
		Bucket: bucket,
	}
}
