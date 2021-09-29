package db

import "C"
import (
	"gopkg.in/mgo.v2"
	"hadouken-base-backend/src/config"
)

type Collection struct {
	db      *Database
	name    string
	Session *mgo.Collection
}

func (c *Collection) Connect() {
	session := *c.db.session.C(c.name)
	c.Session = &session
}
func NewCollection(name string) *Collection {
	dbName, _ := config.Config.String("mongodb::db")
	var c = Collection{
		db:   newDBSession(dbName),
		name: name,
	}
	c.Connect()
	return &c
}
func (c *Collection) Close() {
	DB.Close(c)
}
