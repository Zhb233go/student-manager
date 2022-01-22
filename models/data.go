package models

import (
	. "StudentManager/middleware"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var c = mongo.Collection{}

// Insert 增加一条或多条记录记录
func Insert(ctx context.Context, d bson.M) {
	req, err := c.InsertOne(ctx, d)
	if err != nil {
		Ln.Error(err)
	}
	Ln.Info("insert a single record:", req.InsertedID)
}

// QueryOne 查找一条记录
func QueryOne(ctx context.Context, d bson.D) {
	err := c.FindOne(ctx, d).Decode(&Student{})
	if err != nil {
		Ln.Error(err)
	}
}

func Delete(ctx context.Context, d bson.D) {
	req, err := c.DeleteOne(ctx, d)
	if err != nil {
		Ln.Error(err)
	}
	Ln.Info("delete a single record", req.DeletedCount)
}
