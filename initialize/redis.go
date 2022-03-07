package initialize

import (
	"DRCache/global"
	"fmt"
	"github.com/go-redis/redis"
)

func InitRedis() {
	add := fmt.Sprintf("%s:%d", global.Settings.RedisInfo.Host, global.Settings.RedisInfo.Port)
	// 生成redis客户端
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     add,
		Password: global.Settings.RedisInfo.Password,
		DB:       global.Settings.RedisInfo.DB,
	})
	// 连接redis
	_, err := global.Redis.Ping().Result()
	if err != nil {
		global.Lg.Error("[InitRedis] 连接redis异常")
		global.Lg.Error(global.Settings.RedisInfo.Host)
		global.Lg.Error(err.Error())
	}
}
