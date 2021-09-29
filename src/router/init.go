package router

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"hadouken-base-backend/src/controller"
)

var logger = logs.GetBeeLogger()

func Init() {
	logger.Debug("router init")
	web.AutoRouter(&controller.MainController{})
}
