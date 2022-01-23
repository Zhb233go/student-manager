package models

import (
	. "StudentManager/database"
	. "StudentManager/middleware"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

// InsertOne 增加一条或多条记录记录
func (s *Student) InsertOne(ctx context.Context, d interface{}) {
	req, err := Collection.InsertOne(ctx, d)
	if err != nil {
		Ln.Error(err)
	}
	Ln.Info("insert a single record:", req.InsertedID)
}

// QueryOne 查找一条记录
func (s *Student) QueryOne(ctx context.Context, d interface{}) {
	err := Collection.FindOne(ctx, d).Decode(&Student{})
	if err != nil {
		Ln.Error(err)
	}
	Ln.Info("Queried a single record")
}

func (s *Student) QueryAll(ctx context.Context, d interface{}) []Student {
	a, err := Collection.Find(ctx, d)
	if err != nil {
		Ln.Error(err)
	}
	var results []Student
	if err = a.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results
}

// DeleteOne 删除一条记录
func (s *Student) DeleteOne(ctx context.Context, d bson.D) {
	req, err := Collection.DeleteOne(ctx, d)
	if err != nil {
		Ln.Error(err)
	}
	Ln.Info("delete a single record", req.DeletedCount)
}
