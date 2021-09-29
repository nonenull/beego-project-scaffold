package db

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"hadouken-base-backend/src/config"
)

var logger = logs.GetBeeLogger()

func Init() {
	if DB.baseSession == nil {
		user, _ := config.Config.String("mongodb::user")
		password, _ := config.Config.String("mongodb::password")
		host, _ := config.Config.String("mongodb::host")
		port, _ := config.Config.String("mongodb::port")
		DB.URL = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s", user, password, host, port,
		)
		err := DB.New()
		if err != nil {
			logger.Error("数据库连接发生错误: %s", err)
		}
	}
}
