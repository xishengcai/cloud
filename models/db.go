package models

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/xishengcai/cloud/pkg/db"
)

const (
	Find      = "find"
	Aggregate = "aggregate"
)

func GetCollection(name string) *mongo.Collection {
	return db.DefaultMongoDatabase.Collection(name)
}
