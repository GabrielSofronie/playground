package main

import (
	"gopkg.in/mgo.v2"
	"datastore"
	"datastore/entity"
	"fmt"
)

func main() {
	// Create a Mongo session
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	appCtx := datastore.MongoContext{session.DB("dummy")}
	repo := entity.PageRepo{appCtx.Data.C("content")}

	// First Insert something
	content := `{ "Title": "Delete" }`
	data := &dummyContent{content}
	if err := repo.Insert(data); err != nil {
		panic(err)
	}

	// Then try to Find it
	page, err := repo.Find("How good is that!")
	if err != nil {
		panic(err)
	}

	fmt.Println(page)

	// Delete [One]
	/*
	err := repo.DeleteAll()
	if err != nil {
		panic(err)
	}
	*/
}

type dummyContent struct {
	content string
}

func (dc *dummyContent) Read(data []byte) (int, error) {
	// Do something with content
	//return 0, io.EOF

	for i, b := range []byte(dc.content) {
		data[i] = b
	}

	return len(dc.content), nil
}