package database

import (
	. "StudentManager/middleware"
	_ "StudentManager/utils"
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB, Collection = SetupDB(context.Background())

// SetupDB mongodb数据库初始化
func SetupDB(ctx context.Context) (*mongo.Client, *mongo.Collection) {
	url := viper.GetString("mongodb.url")

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(url)

	// 连接到MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	// 检查连接
	err = client.Ping(ctx, nil)

	// 列出数据库
	data, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		Ln.Error(err)
	}
	Ln.Info(data)
	Ln.Info("Connected to MongoDB!")
	collection := client.Database("StudentManager").Collection("student")
	return client, collection
}
