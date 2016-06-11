package repository

import (
	"testing"
	"datastore"
	"datastore/entities"
)

const (
	path = "/tmp/content/"
	filename = "page.json"
)

func getMockContent() entities.Content {
	return entities.Content{
		Title : "Content Title New",
		Meta : "Content Meta New",
	}
}

// Enforce Interface implement over struct
func TestInterfaceImplement(t *testing.T) {
	var _ datastore.Repository = (*FilesystemPage)(nil)
}

func TestCreate(t *testing.T) {
	page, err := InitFilesystem(path, filename)
	defer page.File.Close()

	if err != nil {
		t.Fatalf("Failed to Init: %v", err)
	}

	if err = page.Create(getMockContent()); err != nil {
		t.Errorf("Failed to Create page: %v", err)
	}
}

func TestRetrieveSingle(t *testing.T) {
	page, err := InitFilesystem(path, filename)
	defer page.File.Close()

	if err != nil {
		t.Fatalf("Failed to Init: %v", err)
	}

	if err = page.Create(getMockContent()); err != nil {
		t.Errorf("Failed to create page: %v", err)
	}

	content, err := page.RetrieveSingle(path + filename)

	if err != nil {
		t.Errorf("Failed to RetrieveSingle page: %v", err)
	}

	t.Logf("RetrieveSingle content: %v", content)
}

func TestRetrieveAll(t *testing.T) {
	page, err := InitFilesystem(path, filename)
	defer page.File.Close()

	if err != nil {
		t.Fatalf("Failed to Init: %v", err)
	}

	if err = page.Create(getMockContent()); err != nil {
		t.Errorf("Failed to create page: %v", err)
	}

	content, err := page.RetrieveAll(path)

	if err != nil {
		t.Errorf("Failed to RetrieveAll pages: %v", err)
	}

	t.Logf("RetrieveAll content: %v", content)
}

func TestDeleteSingle(t *testing.T) {
	page, err := InitFilesystem(path, filename)
	defer page.File.Close()

	if err != nil {
		t.Fatalf("Failed to Init: %v", err)
	}

	if err = page.DeleteSingle(path + filename); err != nil {
		t.Errorf("Failed to DeleteSingle page: %v", err)
	}
}