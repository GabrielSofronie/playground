package entity

import (
	"datastore/models"
	mgo "gopkg.in/mgo.v2"
)

type UserRepo struct {
	Collection *mgo.Collection
}

func (u *UserRepo) Create(user interface{}) error {
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

func (u *UserRepo) Retrieve(id interface{}) (interface{}, error) {
	user := models.User{}
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

func (u *UserRepo) RetrieveBy(field interface{}) (interface{}, error) {
	user := models.User{}
	if err := u.Collection.Find(field).One(&user); err != nil {
		return nil, err
	}
	return user, nil
}

/*
func (u *UserRepo) Update(user interface{}) error {

}
*/

func (u *UserRepo) Delete(user interface{}) error {
	if err := u.Collection.Remove(user); err != nil {
		return err
	}

	return nil
}