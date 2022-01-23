package database

import (
	. "StudentManager/middleware"
	_ "StudentManager/utils"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

var (
	rdb *redis.Client
)

// 初始化连接
func initClient() (err error) {
	options := redis.Options{}
	err = viper.Unmarshal(&options)
	rdb = redis.NewClient(&redis.Options{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
		PoolSize: options.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := rdb.Ping(ctx).Result()
	Ln.Info(result + "redis success")
	return err
}

func V8Example() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		Ln.Fatal("redis: initClient failed: ", err)
		return
	}

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
