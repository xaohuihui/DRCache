package common

import (
	. "DRCache/global"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// author: xaohuihui
// datetime: 2022/2/16 19:28:21
// software: GoLand

type RedisController struct{}

func (r RedisController) SetVal(key string, val interface{}, timeout time.Duration) error {
	err := Redis.Set(key, val, timeout).Err()
	if err != nil {
		return fmt.Errorf("redis set error: %s", err.Error())
	}
	return nil
}

func (r RedisController) GetVal(key string) ([]byte, error) {
	result, err := Redis.Get(key).Bytes()
	if err == redis.Nil {
		return nil, fmt.Errorf("redis key[%s] not exist", key)
	} else if err != nil {
		return nil, fmt.Errorf("redis Get error: %s", err.Error())
	}
	return result, nil
}
