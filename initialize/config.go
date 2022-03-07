package initialize

import (
	"DRCache/config"
	"DRCache/global"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	// 文件的路径设置
	v.SetConfigFile("./settings-dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
	}

	serverConfig := config.ServerConfig{}
	// 给serverConfig 初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		global.Lg.Error("[InitConfig] 初始化config失败")
		global.Lg.Error(err.Error())
		panic(err)
	}
	// 传递全局变量
	global.Settings = serverConfig
	color.Blue("初始化环境变量", global.Settings.LogsAddress)
}
