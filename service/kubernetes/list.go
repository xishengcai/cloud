package kubernetes

import (
	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
)

type List struct {
}

func (l List) Validate() error {
	return nil
}

func (l List) Run() app.ResultRaw {
	data, err := models.GetCluster().Query(nil, models.Find)
	return app.NewServiceResult(data, err)
}
