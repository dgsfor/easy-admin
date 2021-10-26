package conf

import (
	"github.com/joho/godotenv"
	"os"
	"easy-api-template/util"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取配置文件
	// 初始化配置
	SetupConfig()
	//初始化orm
	SetupOrm()
}
