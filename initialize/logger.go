package initialize

import (
	"DRCache/global"
	"DRCache/utils"
	"fmt"
	"go.uber.org/zap"
)

func InitLogger() {
	// 实例化zap 配置
	cfg := zap.NewDevelopmentConfig()
	// 配置日志的输出地址
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", global.Settings.LogsAddress, utils.GetNowFormatTodayTime()),
		"stdout",
	}
	// 创建logger实例
	logg, err := cfg.Build()
	if err != nil {
		fmt.Println(err)
	}
	// 替换zap包中的全局的logger实例， 后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(logg)
	// 注册到全局变量中
	global.Lg = logg
}
