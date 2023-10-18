package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/xishengcai/cloud/pkg/setting"
)

var (
	DefaultMongoClient   *mongo.Client
	DefaultMongoDatabase *mongo.Database
)

func init() {
	cfg := setting.Config.Mongodb
	connectCtx, connectCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer connectCancel()
	uri := fmt.Sprintf("mongodb://%s:%s@%s", cfg.User, cfg.Password, cfg.Host)
	client, err := mongo.Connect(connectCtx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer pingCancel()
	err = client.Ping(pingCtx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	DefaultMongoDatabase = client.Database(setting.Config.Mongodb.Database)
	DefaultMongoClient = client
}
