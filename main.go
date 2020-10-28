package main

import (
	"github.com/spf13/viper"
	"smartlab/conf"
	_ "smartlab/docs"
	"smartlab/router"
	"smartlab/util"

)

// @title smartlab
// @version 1.0
// @descriptionXDU 物理实验计算器 的 Golang 后端

// @contact.name 117503445
// @contact.url http://www.117503445.top
// @contact.email t117503445@gmail.com

// @license.name GNU GPL 3.0
// @license.url https://github.com/TGclub/smartlab_backend_go/blob/main/LICENSE

// @host localhost
// @BasePath /api
// @Schemes http

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	// 从配置文件读取配置
	conf.Init()
	
	// 装载路由
	r := router.NewRouter()

	if err := r.Run(":" + viper.GetString("gin.port")); err != nil {
		util.Log().Panic("router run failed", err)
	}
}
