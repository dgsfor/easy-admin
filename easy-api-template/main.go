package main

import (
	"easy-api-template/conf"
	"easy-api-template/router"
)

func main() {
	// 从配置文件读取配置
	conf.Init()
	// 装载路由
	r := router.NewRouter()
	r.Run(":" + conf.GetConfig("app::port").String())
}
