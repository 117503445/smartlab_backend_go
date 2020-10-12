package main

import (
	"smartlab/conf"
	"smartlab/router"
	"smartlab/util"

	"github.com/spf13/viper"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := router.NewRouter()
	if err := r.Run(":" + viper.GetString("gin.port")); err != nil {
		util.Log().Panic("router run failed", err)
	}
}
