package kubernetes

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"

	"github.com/xishengcai/cloud/models"
	"github.com/xishengcai/cloud/pkg/app"
	"github.com/xishengcai/cloud/pkg/db"
)

var (
	mongoCollection = db.CreateMongoCollection("admin", "cluster")
)

type List struct {
	models.PageInfo
}

func (l List) Validate() error {
	return nil
}

func (l List) Run() app.ResultRaw {
	count, data, err := mongoCollection.SelectPage(context.TODO(), l.getFilter(), nil, l.GetOffset(), l.PageSize)
	return app.NewServiceResultWitRawData(app.ListData{
		Total: count,
		Items: data,
	}, err)
}

func (l List) getFilter() bson.D {
	return bson.D{}
}
