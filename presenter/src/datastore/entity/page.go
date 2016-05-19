package entity

import (
	"datastore/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type PageRepo struct {
	Collection *mgo.Collection
}

func (p *PageRepo) Create(content interface{}) error {
	if err:= p.Collection.Insert(content); err != nil {
		return err
	}
	return nil
}

func (p *PageRepo) Retrieve(id interface{}) (interface{}, error) {
	page := models.Content{}
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

func (p *PageRepo) RetrieveBy(field interface{}) (interface{}, error) {
	content := models.Content{}
	if err := p.Collection.Find(field).One(&content); err != nil {
		return nil, err
	}
	return content, nil
}

func (p *PageRepo) Delete(content interface{}) error {
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