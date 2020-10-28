package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"smartlab/model"
	"smartlab/util"
	"smartlab/docs"
	"github.com/spf13/viper"
)

// Init 初始化配置项
func Init() {



	filepathBase := filepath.Dir(util.GetCurrentPath())
	filePathEnv := filepath.Join(filepathBase, "config.yaml")

	viper.SetConfigFile(filePathEnv)

	if err := viper.ReadInConfig(); err != nil {
		filePathEnv = filepath.Join(filepathBase, "data", "config.yaml")
		viper.SetConfigFile(filePathEnv)

		if err := viper.ReadInConfig(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "load config failed")
			panic(err)
		}
	}

	docs.SwaggerInfo.Host = viper.GetString("swagger.host")
	docs.SwaggerInfo.Schemes=viper.GetStringSlice("swagger.scheme")


	// 设置日志级别
	util.BuildLogger(viper.GetString("log.level"))

	// 连接数据库
	model.InitDatabase()
}
