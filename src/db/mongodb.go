package db

import (
	"gopkg.in/mgo.v2"
	"hadouken-base-backend/src/config"
)

type MongoDB struct {
	baseSession *mgo.Session
	queue       chan int
	URL         string
	Open        int
}

var DB MongoDB

func (mongodb *MongoDB) New() error {
	maxPool, err := config.Config.Int("mongodb::pool.size")
	if err != nil {
		logger.Error("读取mongodb pool.size 发生错误: %s", err)
	}
	mongodb.queue = make(chan int, maxPool)
	for i := 0; i < maxPool; i = i + 1 {
		mongodb.queue <- 1
	}
	mongodb.Open = 0
	mongodb.baseSession, err = mgo.Dial(mongodb.URL)
	return err
}
func (mongodb *MongoDB) Session() *mgo.Session {
	<-mongodb.queue
	mongodb.Open++
	return mongodb.baseSession.Copy()
}
func (mongodb *MongoDB) Close(c *Collection) {
	c.db.s.Close()
	mongodb.queue <- 1
	mongodb.Open--
}
