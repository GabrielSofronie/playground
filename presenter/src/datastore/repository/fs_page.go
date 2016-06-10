package repository

import (
	_ "datastore"
//	"datastore/entities"
	"encoding/json"
	"errors"
//	"io/ioutil"
	"os"
	"io/ioutil"
	"datastore/entities"
)

// This implementation creates some sample files inside /tmp
// for the sole purpose to be used as a DB replacement.
// There is no proper error-checking in place nor specific cases
// are considered; a better implementation should
// consider adding proper validation, directory creation etc.

type Repository interface {
	Create(interface{}) error
	RetrieveSingle(interface{}) (interface{}, error)
	RetrieveAll(interface{}) (interface{}, error)
	DeleteSingle(interface{}) error
//	Drop() error
}

type FilesystemPage struct {
	*os.File
}

func InitFilesystem(filename string) (*FilesystemPage, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	p := &FilesystemPage{f}

	return p, nil
}

func (p *FilesystemPage) Create(content interface{}) error {
	page, err := json.Marshal(content)

	if err != nil {
		return err
	}

	if _, err = p.File.Write(page); err != nil {
		return err
	}

	return nil
}

func (p *FilesystemPage) RetrieveSingle(id interface{}) (interface{}, error) {
	if p.File.Name() != id.(string) {
		return nil, errors.New("File name not matching!")
	}

	content, err := ioutil.ReadFile(p.File.Name())
	if err != nil {
		return nil, err
	}

	page := entities.Content{}
	if err = json.Unmarshal(content, &page); err != nil {
		return nil, err
	}

	return page, nil
}

func (p *FilesystemPage) RetrieveAll(collection interface{}) (interface{}, error) {
	fileList, err := ioutil.ReadDir(collection.(string))

	if err != nil {
		return nil, err
	}

	var content []entities.Content

	for _, file := range fileList {
		c, err := ioutil.ReadFile(collection.(string) + file.Name())
		if err != nil {
			return nil, err
		}

		page := entities.Content{}
		if err = json.Unmarshal(c, &page); err != nil {
			return nil, err
		}
		content = append(content, page)
	}

	return content, nil
}

func (p *FilesystemPage) DeleteSingle(id interface{}) error {
	if p.File.Name() != id.(string) {
		return errors.New("File name not matching!")
	}

	return os.Remove(p.File.Name())
}

/*
func (p *FilesystemPage) Drop() error {
	return os.Remove(p.Filename)
}
*/