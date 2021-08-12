package sync

import (
	cr20181201 "github.com/alibabacloud-go/cr-20181201/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
)

// AliCloud get aliCloud image register
type AliCloud struct {
	region          string
	accessKeyID     string
	accessKeySecret string
	*cr20181201.Client
}

// NewAliImageRegistry return AliCloud
func NewAliImageRegistry(region, accessKeyID, accessKeySecret string) AliCloud {
	return AliCloud{
		region:          region,
		accessKeySecret: accessKeySecret,
		accessKeyID:     accessKeyID,
	}
}

// setClient enable AliCloud has client function
func (a *AliCloud) setClient() error {
	config := &openapi.Config{
		AccessKeyId:     &a.accessKeyID,
		AccessKeySecret: &a.accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = &a.region
	client, err := cr20181201.NewClient(config)
	a.Client = client
	return err
}
