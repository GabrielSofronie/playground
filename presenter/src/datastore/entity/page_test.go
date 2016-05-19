package entity

import (
	db "datastore"
	"datastore/models"
//	"testing"
	"time"
	"testing"
)

/*
func TestGetTitle(t *testing.T) {
	testCases := []struct {
		p Page
		expectedTitle string
	}{
		{
			p: Page{
				Content: models.Content{
					Title: "Page title",
				},
			},
			expectedTitle: "Page title",
		},
		{
			p: Page{
				Content: models.Content{
					Title: "",
				},
			},
			expectedTitle: "",
		},
	}

	for _, c := range testCases {
		if c.expectedTitle != c.p.GetTitle() {
			t.Errorf("Expected title: %s but got %s", c.expectedTitle, c.p.GetTitle())
		}
	}
}

func TestGetPage(t *testing.T) {
	page, err := GetPage()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if page.Title == "" {
		t.Error("Empty title")
	}
}
*/

type mockDataStore struct {}

func (mds *mockDataStore)Find(interface{}) (interface{}, error) {
	return models.Content{
		Title: "Mock page title",
		/*
		Author: models.User{
			Name: "Author",
			Age: 25,
			Registered: time.Now(),
			Updated: time.Now(),
		},
		*/
		Created: time.Now(),
		Updated: time.Now(),
	}, nil
}
/*
func TestFind(t *testing.T) {
	mds := mockDataStore{}
	p, err := mds.Find("name")
	if err != nil {
		t.Errorf("There was an error searching for Page: %v", err)
	}

	t.Log(p)
}
*/

// Enforce Interface implement over struct
func TestInterfaceImplement(t *testing.T) {
	var _ db.Repository = (*PageRepo)(nil)
}