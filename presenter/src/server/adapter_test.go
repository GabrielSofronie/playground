package server

import (
	//"testing"
	/*
	"net/http"
	"fmt"
	"net/http/httptest"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/context"
	*/
)

/*
func mockHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mock Handler")
	})
}

func mockAdapter(rec *httptest.ResponseRecorder, req *http.Request) HandlerAdapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(rec, req)
		})
	}
}

func TestAdapt(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Errorf("Request failed: %v", err)
	}
	HandleAdapt(mockHandler(), mockAdapter(httptest.NewRecorder(), req))
}

func TestWithDb(t *testing.T) {
	db, err := mgo.Dial("localhost")
	if err != nil {
		t.Fatalf("Dial error: %v", err)
	}
	defer db.Close()

	h := HandleAdapt(mockHandler(), AdaptDB(db))

	http.Handle("/user", context.ClearHandler(h))
}
*/

/*

curl -H "Content-Type: application/json" -X POST -d '{"Title":"Super Awesome", "Meta":"Rock"}' http://localhost:8080/content/fs/file-one

*/