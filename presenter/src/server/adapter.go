package server

import (
	"net/http"
//	"gopkg.in/mgo.v2"
	"github.com/gorilla/context"
	"datastore/repository"
	"datastore/entities"
	"encoding/json"
	"path"

//	"log"
//	"fmt"
	"text/template"
	//"datastore"
)

const (
	ctxDatastore = "datastore"
	tplFile = "../resources/views/master"
	tplExt = ".got"
	fsPath = "/tmp/"
)

var filename string

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
func GetFSHandler(urlPath string) http.Handler {
	//fs := repository.FilesystemPage{"/tmp/page.json"}
	filename = path.Base(urlPath)

	conf := map[string]string{
		"repository" : "filesystem",
		"path" : fsPath,
		"filename" : filename,
	}

	fs, err := repository.GetRepository(conf)

	if err != nil {
		panic(err)
	}

	return HandleAdapt(http.HandlerFunc(handleFile), adaptDatastore(fs))
}

func handleFile(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		viewFile(w, r)
	case "POST":
		createFile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
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
	fs := context.Get(r, ctxDatastore).(*repository.FilesystemPage)

	var content entities.Content

	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		panic(err)
	}

	if err := fs.Create(content); err != nil {
		panic(err)
	}
}

func viewFile(w http.ResponseWriter, r *http.Request) {
	fs := context.Get(r, ctxDatastore).(*repository.FilesystemPage)

	page, err := fs.RetrieveSingle(fsPath + filename)

	if err != nil {
		panic(err)
	}

	renderContent(w, tplFile, page)
}

func renderContent(w http.ResponseWriter, tpl string, content interface{}) {
	t, _ := template.ParseFiles(tpl + tplExt)
	t.Execute(w, content)
}