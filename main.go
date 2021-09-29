package main

import (
	_ "github.com/beego/beego/v2/adapter/session/redis"
	_ "github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"gopkg.in/mgo.v2/bson"
	"hadouken-base-backend/src/config"
	"hadouken-base-backend/src/db"
	"hadouken-base-backend/src/router"
)

var logger = logs.GetBeeLogger()

type Person struct {
	Name  string
	Phone string
}

func init() {
	// logs
	logs.SetLevel(logs.LevelDebug)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	config.Init()
	db.Init()
	// beego
	web.SetStaticPath("/assets", "assets")
	router.Init()
}

func main() {
	c := db.NewCollection("fuck")
	var err error
	err = c.Session.Insert(&Person{"Ale", "111111"}, &Person{"Cla", "222222222"})

	var result []Person
	err = c.Session.Find(bson.M{"name": "Ale"}).All(&result)
	indexes, _ := c.Session.Indexes()
	logger.Debug("indexs === %+v", indexes)
	if err != nil {
		logger.Error("err ==== %s", err)
	}
	logger.Debug("data === %s %s", result, err)
	web.Run()
}
