package repository

import (
	"datastore/entities"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

/*
*** This is an EntityGatewayImplementer ***
It uses a Database API (mgo - in this case) and provides the query-methods to work with the DB
e.g. RetrieveBy(...)
These ones can be stubbed to simulate a response from DB, making their use (and unit-testing) easier
*/

/*
Define a MongoPage which implements Repository interface;
Can add another implementation for a FilesystemPage which satisfies the same interface.
Or even a MockUser etc.
*/

type MongoPage struct {
	Collection *mgo.Collection
}

func (p *MongoPage) Create(content interface{}) error {
	if err:= p.Collection.Insert(content); err != nil {
		return err
	}
	return nil
}

func (p *MongoPage) Retrieve(id interface{}) (interface{}, error) {
	page := entities.Content{}
	// Use type assertion to transform id -> string
	if str, err := id.(string); err {
		err := p.Collection.FindId(bson.ObjectIdHex(str)).One(&page)
		if err != nil {
			return page, err
		}
	} else {
		return page, errors.New("Provided Page ID cannot be converted to string!")
	}

	return page, nil
}

func (p *MongoPage) RetrieveBy(field interface{}) (interface{}, error) {
	content := entities.Content{}
	if err := p.Collection.Find(field).One(&content); err != nil {
		return nil, err
	}
	return content, nil
}

func (p *MongoPage) Delete(content interface{}) error {
	if err := p.Collection.Remove(content); err != nil {
		return err
	}

	return nil
}

/*
func (pr *PageRepo) Find(id interface{}) (interface{}, error) {
	page := models.Content{}
	// Use type assertion to transform id -> string
	if str, err := id.(string); err {
		err := pr.Collection.Find(bson.M{"title" : str}).One(&page)
		if err != nil {
			return page, err
		}
	} else {
		return page, errors.New("Provided Page TITLE cannot be converted to string!")
	}

	return page, nil
}

// Mostly duplicate of Find

func (pr *PageRepo) Insert(data io.Reader) error {
	page := models.Content{}
	if err := json.NewDecoder(data).Decode(&page); err != nil {
		return err
	}

	if err:= pr.Collection.Insert(page); err != nil {
		return err
	}
	return nil
}

func (pr *PageRepo) DeleteById(id interface{}) error {
	if err := pr.Collection.Remove(id); err != nil {
		return err
	}

	return nil
}
*/