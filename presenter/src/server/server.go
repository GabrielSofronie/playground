package server

import (
	"log"
	"net/http"
	"github.com/gorilla/context"

	/*
	"gopkg.in/mgo.v2"
	"datastore"
	"datastore/repository"
	"datastore/entities"
	"fmt"
	"github.com/gorilla/context"
	*/
)

// Start a basic server
func Start() {
	//databaseHandler := GetDBHandler()

	urlPath := "/content/fs/file-one"
	fileHandler := GetFSHandler(urlPath)

	// sample path
	// SHOULD USE Gorilla Mux and RegEx Paths
	http.Handle(urlPath, context.ClearHandler(fileHandler))

	// start a web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Routing to handlers and use packages

// Return some handlers
// Uses an Adapter
// NOTE: http.HandlerFunc IS ALSO an ADAPTER -- check the docs
/*
func GetDBHandler() http.Handler {
	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	// FIND HOW defer works
	// Uncommenting the following line will panic (session already closed)
	//defer db.Close()

	return HandleAdapt(http.HandlerFunc(getActionForMethod), AdaptDB(db))
}

// This should be moved and refactored
func getActionForMethod(w http.ResponseWriter, r *http.Request) {
	//switch r.Method {
	//case "GET" :
		handleRetrieve(w, r)
	//default :
	//	http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	//}
}

func handleRetrieve(w http.ResponseWriter, r *http.Request) {
	// Get DB session from context value
	session := context.Get(r, "datastore").(*mgo.Session)
	// IMPORTANT: DB name should be stored in a Config file
	appCtx := datastore.MongoContext{session.DB("dummy")}

	// Dummy test a User
	repo := repository.MongoUser{appCtx.C("users")}

	// Dummy insert
	user := entities.User{
		Name: "Architect",
	}
	if err := repo.Create(user); err != nil {
		log.Printf("User create error: %v", err)
	}
	// Display Dummy user
	u, err := repo.RetrieveBy(user)
	if err != nil {
		log.Printf("Retrieve error:\n%v\n%v", user, err)
	}

	log.Println(u)
}
*/