package server

import (
	"net/http"
//	"gopkg.in/mgo.v2"
	"github.com/gorilla/context"
	"datastore/repository"
	"datastore/entities"

//	"log"
//	"fmt"
	"text/template"
)

const (
	ctxDatastore = "datastore"
	tplExt = ".got"
)

type HandlerAdapter func(http.Handler) http.Handler

func HandleAdapt(handle http.Handler, adapters ...HandlerAdapter) http.Handler {
	for _, adapter := range adapters {
		handle = adapter(handle)
	}

	return handle
}

// STORAGE HANDLERS

// REFACTOR and MOVE to own file
/*
func AdaptDB(db *mgo.Session) Adapter {
	// return the Adapter
	return func(h http.Handler) http.Handler {
		// the adapter (when called) should return a new handler
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// copy the database session
			dbsession := db.Copy()
			defer dbsession.Close() // clean up
			// save it in the mux context
			context.Set(r, ctxDatastore, dbsession)
			// pass execution to the original handler
			h.ServeHTTP(w, r)
		})
	}
}
*/

// !!!This is an experimental Filesystem Handler
func GetFSHandler() http.Handler {
	fs := repository.FilesystemPage{"/tmp/page.json"}

	return HandleAdapt(http.HandlerFunc(handleFile), adaptDatastore(fs))
}

func handleFile(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		createFile(w, r)
		viewFile(w, r)
	case "POST":
		createFile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	/*
	fs := context.Get(r, "filesystem").(repository.FilesystemPage)

	// Dummy insert
	content := entities.Content{
		Title: "Architect",
	}
	// Test by creating a new file
	log.Printf("Create on the FS: %v", fs.Create(content))

	// Try to retrieve and display the content
	page, _ := fs.Retrieve("")
	fmt.Fprintf(w, "Tadaaaa: %v", page)
	*/
}

// Can use Filesystem instead of DB-store
func adaptDatastore(ds interface{}) HandlerAdapter {
	return func (h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, ctxDatastore, ds)
			h.ServeHTTP(w, r)
		})
	}
}

// ACTION HANDLERS
func createFile(w http.ResponseWriter, r *http.Request) {
	fs := context.Get(r, ctxDatastore).(repository.FilesystemPage)

	content := entities.Content{
		Title: "Architect",
	}

	if err := fs.Create(content); err != nil {
		panic(err)
	}
}

func viewFile(w http.ResponseWriter, r *http.Request) {
	fs := context.Get(r, ctxDatastore).(repository.FilesystemPage)

	page, err := fs.Retrieve("")

	if err != nil {
		panic(err)
	}

	tplFile := "../resources/views/master"
	renderContent(w, tplFile, page)
}

func renderContent(w http.ResponseWriter, tpl string, content interface{}) {
	t, _ := template.ParseFiles(tpl + tplExt)
	t.Execute(w, content)
}