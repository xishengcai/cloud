package models

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/xishengcai/cloud/pkg/common"
)

// Cluster kubernetes cluster meta data
type Cluster struct {
	ID          string `bson:"id" json:"-"`
	Name        string `bson:"name" json:"name" validate:"required" default:"test"`
	Master      []Host `bson:"master" json:"master" validate:"required"`
	NetWorkPlug string `bson:"networkPlug" json:"networkPlug" default:"cilium"`
	// registry.aliyuncs.com/google_containers， k8s.gcr.io
	Registry             string `bson:"registry" json:"registry" default:"registry.aliyuncs.com/google_containers"`
	Version              string `bson:"version" json:"version" default:"1.22.15"`
	ControlPlaneEndpoint string `bson:"controlPlaneEndpoint" json:"controlPlaneEndpoint" validate:"required"`
	PodCidr              string `bson:"podCidr" json:"podCidr" default:"10.244.0.0/16"`
	ServiceCidr          string `bson:"serviceCidr" json:"serviceCidr" default:"10.96.0.0/16"`
	SlaveNode            []Host `bson:"slaveNode" json:"slaveNode"`
}

// KubernetesSlave k8s slave node
type KubernetesSlave struct {
	Version string `form:"version"`
	Nodes   []Host `form:"nodes"`
	Master  Host   `form:"master"`
}

// Version cluster version
type Version struct {
	Version string `form:"version" default:"1.22.15"`
}

func GetCluster() *Cluster {
	return &Cluster{}
}

func (c *Cluster) Query(filter interface{}, options string) ([]Cluster, error) {
	collection := GetCollection("cluster")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	clusterList := make([]Cluster, 0)
	cur := &mongo.Cursor{}
	var err error
	switch options {
	case Find:
		cur, err = collection.Find(ctx, filter)
		if err != nil {
			return nil, err
		}
	case Aggregate:
		cur, err = collection.Aggregate(ctx, filter)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("mongo query option not support")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var cluster Cluster
		err := cur.Decode(&cluster)
		if err != nil {
			return nil, err
		}
		clusterList = append(clusterList, cluster)
	}
	return clusterList, nil
}

func (c *Cluster) Insert() error {
	c.ID = common.GetUUID()
	collection := GetCollection("cluster")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := collection.InsertOne(ctx, c)
	if err != nil {
		return err
	}
	return nil
}
