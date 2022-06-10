package global

// author: xaohuihui
// datetime: 2022/3/7 15:16:10
// software: GoLand

import (
	"DRCache/config"
	"DRCache/utils"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Settings       config.ServerConfig
	Lg             *zap.Logger
	Trans          ut.Translator
	DB             *gorm.DB
	Redis          *redis.Client
	ConsistentHash *utils.Consistent
)
