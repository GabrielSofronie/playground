package repository

import (
	"datastore/entities"
	mgo "gopkg.in/mgo.v2"
)

/*
Define a MongoUser which implements Repository interface;
Can add another implementation for a FilesystemUser which satisfies the same interface.
Or even a MockUser etc.
*/

type MongoUser struct {
	Collection *mgo.Collection
}

func (u *MongoUser) Create(user interface{}) error {
	if err:= u.Collection.Insert(user); err != nil {
		return err
	}
	return nil
}

/*
func (u *UserRepo) InsertFromSource(data io.Reader) error {
	user := models.User{}
	if err := json.NewDecoder(data).Decode(&user); err != nil {
		return err
	}

	if err:= u.Collection.Insert(user); err != nil {
		return err
	}
	return nil
}
*/

func (u *MongoUser) Retrieve(id interface{}) (interface{}, error) {
	user := entities.User{}
	// Use type assertion to transform id -> string

	u.Collection.Find(id).One(&user)

	/*
	if str, err := id.(string); err {
		//err := pr.Collection.FindId(bson.ObjectIdHex(str)).One(&page)
		err := pr.Collection.Find(bson.M{"name" : str}).One(&user)
		if err != nil {
			return user, err
		}
	} else {
		return user, errors.New("Provided User NAME cannot be converted to string!")
	}
	*/

	return user, nil
}

func (u *MongoUser) RetrieveBy(field interface{}) (interface{}, error) {
	user := entities.User{}
	if err := u.Collection.Find(field).One(&user); err != nil {
		return nil, err
	}
	return user, nil
}

/*
func (u *UserRepo) Update(user interface{}) error {

}
*/

func (u *MongoUser) Delete(user interface{}) error {
	if err := u.Collection.Remove(user); err != nil {
		return err
	}

	return nil
}