package models

import (
	. "StudentManager/database"
	. "StudentManager/middleware"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// QueryAll 查找多条记录
func (s *Student) QueryAll(ctx context.Context, d interface{}) (ss []*Student) {
	// 查询多个
	// 将选项传递给Find()
	findOptions := options.Find()
	findOptions.SetLimit(500)

	// 定义一个切片用来存储查询结果
	var results []*Student

	// 把bson.D{{}}作为一个filter来匹配所有文档
	cur, err := Collection.Find(ctx, bson.D{{}}, findOptions)

	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(ctx) {

		// 创建一个值，将单个文档解码为该值
		var elem Student
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err = cur.Err(); err != nil {
		Ln.Error(err)
	}

	// 完成后关闭游标
	cur.Close(ctx)
	Ln.Info("Found multiple documents (array of pointers)")
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
