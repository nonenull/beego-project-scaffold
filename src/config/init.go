package config

import (
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"os"
)

var Config config.Configer
var Path = "config/config.ini"
var logger = logs.GetBeeLogger()

func InitMongoDBConf() {

}

func InitBeegoConf() {
	// 初始化 session
	web.BConfig.WebConfig.Session.SessionOn = true
	web.BConfig.WebConfig.Session.SessionProvider = "redis"
	redisHost, _ := Config.String("redis::host")
	redisPort, _ := Config.String("redis::port")
	redisPassword, _ := Config.String("redis::password")
	redisPoolSize, _ := Config.String("redis::pool.size")
	redisDb, _ := Config.String("session::db")
	if redisPassword == "" {
		web.BConfig.WebConfig.Session.SessionProviderConfig = fmt.Sprintf("%s:%s,%s,%s,", redisHost, redisPort, redisPoolSize, redisDb)
	} else {
		web.BConfig.WebConfig.Session.SessionProviderConfig = fmt.Sprintf("%s:%s,%s,%s,%s", redisHost, redisPort, redisPoolSize, redisPassword, redisDb)
	}
}

func Init() {
	var err error
	Config, err = config.NewConfig("ini", Path)
	if err != nil {
		logger.Error("读取配置文件 [%s] 发生错误 %s", Path, err)
		os.Exit(1)
	}
	InitBeegoConf()
	InitMongoDBConf()
}
