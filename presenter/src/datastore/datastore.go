package datastore

import (
	"io"
	"gopkg.in/mgo.v2"
)

type Datastore interface {
	//Find()
	Find(interface{}) (interface{}, error)
	Insert(io.Reader) error
	//SelectAll() (interface{}, error)
}

type MongoContext struct {
	Data *mgo.Database
}