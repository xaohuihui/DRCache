package common

import (
	"DRCache/DRCachePkg/config"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

// author: xaohuihui
// datetime: 2022/2/17 11:26:16
// software: GoLand

// 全局业务变量
var (
	BusErrorInstance = &BusError{}
)

// 全局配置变量
var (
	RedisConfigInstance  = &config.RedisConfig{}
	LogrusConfigInstance = &config.LogrusConfig{}
)

// 全局客户端变量
var (
	RedisClient *redis.Client
	LoggerClient = logrus.New()
)

// 全局常量
const (
	ConfigPath = "./DRCachePkg/config/" //配置文件目录
	StaticPath = ""                     //静态资源文件目录
	EnvDev     = "dev"
	EnvLocal   = "local"
	EnvProd    = "prod"
	EnvPrepub  = "prod"
)
