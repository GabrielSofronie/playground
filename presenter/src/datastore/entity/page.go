package entity

import (
	//db "datastore"
	"datastore/models"
	"errors"
	"encoding/json"
	"io"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Query interface {
	GetTitle() string
}

type Page struct {
	Content models.Content
}

func (p Page)GetTitle() string {
	return p.Content.Title
}

func GetPage() (models.Content, error) {
	return models.Content{}, nil
}

// -------------------------------------

/*
type Filesystem struct {}

func (fs *Filesystem)SelectAll() (interface{}, error) {
	return models.Content{}, nil
}

func tst() {
	var _ db.Datastore = (*Filesystem)(nil)
}
*/

type PageRepo struct {
	Collection *mgo.Collection
}

func (pr *PageRepo) Find(id interface{}) (interface{}, error) {
	page := models.Content{}
	// Use type assertion to transform id -> string
	if str, err := id.(string); err {
		//err := pr.Collection.FindId(bson.ObjectIdHex(str)).One(&page)
		err := pr.Collection.Find(bson.M{"title" : str}).One(&page)
		if err != nil {
			return page, err
		}
	} else {
		return page, errors.New("Provided Page ID cannot be converted to string!")
	}

	return page, nil
}

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

func (pr *PageRepo) DeleteAll(id interface{}) error {
		if err := pr.Collection.Remove(id); err != nil {
		return err
	}

	return nil
}