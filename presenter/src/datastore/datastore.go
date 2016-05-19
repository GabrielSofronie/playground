package datastore

import "gopkg.in/mgo.v2"

// CRUD operations over Repository
type Repository interface {
	Create(interface{}) error
	Retrieve(interface{}) (interface{}, error)
	//Update(...)
	Delete(interface{}) error

	//Find()
	//Find(interface{}) (interface{}, error)
	//Insert(io.Reader) error
	//SelectAll() (interface{}, error)
}

type MongoContext struct {
	*mgo.Database
}