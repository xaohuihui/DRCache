package main

import (
	"DRCache/app/queryCache"
	"DRCache/global"
	"DRCache/initialize"
	"flag"
	"fmt"
)

// author: xaohuihui
// datetime: 2022/2/16 18:27:03
// software: GoLand

func startAPIServer() {
	// 初始化yaml配置
	initialize.InitConfig()

	// 加载多个APP的路由配置
	initialize.Include(queryCache.Routers)
	// 初始化路由
	r := initialize.InitRouters()

	// 初始化日志信息
	initialize.InitLogger()

	// 初始化PG
	//initialize.InitPGDB()

	// 初始化redis
	initialize.InitRedis()

	if err := r.Run(fmt.Sprintf(":%d", global.Settings.Port)); err != nil {
		fmt.Printf("startup service failed, err: %v\n", err)
	}
}

func startRPCServer(addr string) {
	// 初始化yaml配置
	initialize.InitConfig()

	// 初始化日志信息
	initialize.InitLogger()

	// 初始化redis
	initialize.InitRedis()

	initialize.InitRPCServer("tcp", addr)
}

func main() {
	var addr string
	var api bool
	flag.StringVar(&addr, "addr", ":8081", "RPC server addr")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	if api {
		go startAPIServer()
	}
	startRPCServer(addr)
}
