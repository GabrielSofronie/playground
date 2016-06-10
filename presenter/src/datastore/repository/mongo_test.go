package repository

import "gopkg.in/mgo.v2"

/*
import (
	"testing"
	"gopkg.in/mgo.v2"
	"datastore"
	"datastore/entities"
	"time"
	"datastore"
	"datastore/entities"
)
*/

type dummyContent struct {
	content string
}

// implement io.Read
func (dc *dummyContent) Read(data []byte) (int, error) {
	// Do something with content
	//return 0, io.EOF

	for i, b := range []byte(dc.content) {
		data[i] = b
	}

	return len(dc.content), nil
}

func getMongoSession() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	return session, nil
}

/*
func TestMongoPageCrud(t *testing.T) {
	_, err := getMongoSession()

	if err != nil {
		t.Errorf("Cannot get Mongo Session: %v", err)
	}

	appCtx := datastore.MongoContext{session.DB("dummy")}
	repo := PageRepo{appCtx.C("content")}

	// First Insert something
	content := entities.Content{
		Title: "Awesome",
	}
	if err := repo.Create(content); err != nil {
		t.Errorf("Create error: %v", err)
	}

	// Then try to Find it
	content = entities.Content{
		Title: "Awesome",
	}
	page, err := repo.RetrieveBy(content)
	if err != nil {
		t.Errorf("Retrieve error: %v", err)
	}

	t.Logf("[%v] Model: %v \n%v", time.Now(), content, page)

	// Delete [One]
	content = entities.Content{
		Title: "Awesome",
	}
	err = repo.Delete(content)

	if err != nil {
		t.Errorf("Error Deleting content: \n%v\n%v", content, err)
	}
}

/*
func TestMongoUserCrud(t *testing.T) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		t.Fatalf("Dial error: %v", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	appCtx := datastore.MongoContext{session.DB("dummy")}
	repo := UserRepo{appCtx.C("users")}

	// First Insert something
	user := entities.User{
		Name: "14th One",
		//Registered: time.Now(),
	}
	t.Log(user)
	//data := &dummyContent{user}
	if err := repo.Create(user); err != nil {
		t.Errorf("Create error: %v", err)
	}

	// Then try to Find it
	user = entities.User{
		Name: "14th One",
	}
	u, err := repo.RetrieveBy(user)
	if err != nil {
		t.Errorf("Retrieve error:\n%v\n%v", user, err)
	}

	t.Logf("[%v] Model: %v \n%v", time.Now(), user, u)

	// Delete [One]
	user = entities.User{
		Name: "14th One",
	}
	err = repo.Delete(user)
	if err != nil {
		t.Errorf("Error Deleting user: %v", err)
	}
}
*/
