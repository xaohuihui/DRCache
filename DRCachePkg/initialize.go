package DRCachePkg

// author: songyanhui
// datetime: 2022/2/16 19:49:58
// software: GoLand

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var RedisClient *redis.Client

var ctx = context.Background()

func connectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	duration, err := time.ParseDuration("1000")
	ctx, CancelFunc := context.WithTimeout(context.Background(), duration)
	defer CancelFunc()
	result, err := RedisClient.Ping(ctx).Result()
	fmt.Println("redis: " + result)
	fmt.Println(err)
}
