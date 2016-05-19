package entity

import (
	"testing"
	"gopkg.in/mgo.v2"
	"datastore"
	"datastore/models"
	"time"
)

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

func TestMongoDial(t *testing.T) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		t.Fatalf("Dial error: %v", err)
	}
	defer session.Close()
}

func TestMongoPageCrud(t *testing.T) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		t.Fatalf("Dial error: %v", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	appCtx := datastore.MongoContext{session.DB("dummy")}
	repo := PageRepo{appCtx.C("content")}

	// First Insert something
	content := models.Content{
		Title: "Awesome",
	}
	if err := repo.Create(content); err != nil {
		t.Errorf("Create error: %v", err)
	}

	// Then try to Find it
	content = models.Content{
		Title: "Awesome",
	}
	page, err := repo.RetrieveBy(content)
	if err != nil {
		t.Errorf("Retrieve error: %v", err)
	}

	t.Logf("[%v] Model: %v \n%v", time.Now(), content, page)

	// Delete [One]
	content = models.Content{
		Title: "Awesome",
	}
	err = repo.Delete(content)

	if err != nil {
		t.Errorf("Error Deleting content: \n%v\n%v", content, err)
	}
}

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
	user := models.User{
		Name: "14th One",
		//Registered: time.Now(),
	}
	t.Log(user)
	//data := &dummyContent{user}
	if err := repo.Create(user); err != nil {
		t.Errorf("Create error: %v", err)
	}

	// Then try to Find it
	user = models.User{
		Name: "14th One",
	}
	u, err := repo.RetrieveBy(user)
	if err != nil {
		t.Errorf("Retrieve error:\n%v\n%v", user, err)
	}

	t.Logf("[%v] Model: %v \n%v", time.Now(), user, u)

	// Delete [One]
	user = models.User{
		Name: "14th One",
	}
	err = repo.Delete(user)
	if err != nil {
		t.Errorf("Error Deleting user: %v", err)
	}
}
