package database

import (
	. "StudentManager/middleware"
	_ "StudentManager/utils"
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var DB = SetupDB(context.Background())

// SetupDB mongodb数据库初始化
func SetupDB(ctx context.Context) *mongo.Client {
	url := viper.GetString("mongodb.url")
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(url)

	// 连接到MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	Ln.Println("Connected to MongoDB!")
	return client
}
