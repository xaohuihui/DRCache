package DRCachePkg

// author: songyanhui
// datetime: 2022/2/16 19:49:58
// software: GoLand

import (
	. "DRCache/DRCachePkg/common"
	"context"
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type App struct {
	Env string `ini:"env"`
}

func (app *App) loadConfig() {
	iniPath := ConfigPath + app.Env + ".ini"
	cfg, err := ini.Load(iniPath)
	BusErrorInstance.ThrowError(err)
	err = cfg.Section("app").MapTo(app)
	// 加载redis配置
	err = cfg.Section("redis").MapTo(RedisConfigInstance)
	BusErrorInstance.ThrowError(err)
}

func connectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisConfigInstance.Host + ":" + RedisConfigInstance.Port,
		Password: RedisConfigInstance.Password,  // no password set
		DB:       RedisConfigInstance.DefaultDB, // use default DB
		PoolSize: RedisConfigInstance.PoolSize,
	})
	duration, err := time.ParseDuration(RedisConfigInstance.TimeOut)
	_, CancelFunc := context.WithTimeout(context.Background(), duration)
	defer CancelFunc()
	result, err := RedisClient.Ping().Result()
	fmt.Println("redis: " + result)
	fmt.Println(err)
}

// 设置全局日志Logger
func setLoggerInstance() {
	// 设置日志等级
	var level logrus.Level
	err := level.UnmarshalText([]byte(LogrusConfigInstance.Level))
	BusErrorInstance.ThrowError(err)
	LoggerClient.SetLevel(level)
	// 设置日志格式
	switch LogrusConfigInstance.Formatter {
	case "json":
		LoggerClient.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		LoggerClient.SetFormatter(&logrus.TextFormatter{})
	case "customize":
		LoggerClient.SetFormatter(&CustomizeFormat{})
	default:
		BusErrorInstance.ThrowError(errors.New("log formatter must json|text|customize"))
	}
	// 打开日志记录的行数； true开启 false关闭 默认false
	if LogrusConfigInstance.ReportCaller {
		LoggerClient.SetReportCaller(LogrusConfigInstance.ReportCaller)
	}
	// 设置日志输出方式
	switch LogrusConfigInstance.OutputType {
	case "1":
		LoggerClient.SetOutput(os.Stdout)
	case "2":
		Log2FileByClass()
	default:
		LoggerClient.SetOutput(os.Stdout)
	}
}
