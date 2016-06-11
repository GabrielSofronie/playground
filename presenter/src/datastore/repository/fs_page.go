package repository

import (
	"encoding/json"
	"errors"
	"os"
	"io/ioutil"
	"datastore/entities"
	"fmt"
	"datastore"
)


func FilesystemPageRepository(conf map[string]string) (datastore.Repository, error) {
	path := conf["path"]
	filename := conf["filename"]

	if path == "" || filename == "" {
		return nil, errors.New(fmt.Sprintf("Missing configuration data:\npath: %v\nfilename: %v", path, filename))
	}

	return InitFilesystem(path, filename)
}

func MongoPageRepository(conf map[string]string) (datastore.Repository, error) {
	return nil, nil
}

//----------------------------------------------------
//-----   Maybe this mapping can be avoided and  -----
//-----         a simpler Get can be used        -----
//----------------------------------------------------
var repoFactories = make(map[string]datastore.RepositoryFactory)

func Register(name string, factory datastore.RepositoryFactory) error {
	if factory == nil {
		return errors.New("No valid factory provided")
	}

	repoFactories[name] = factory

	return nil
}

func initialize() {
	Register("filesystem", FilesystemPageRepository)
}
//----------------------------------------------------

func GetRepository(conf map[string]string) (datastore.Repository, error) {
	initialize()

	repositoryName := conf["repository"]

	if repositoryName == "" {
		return nil, errors.New("Repository not found! Empty values are not allowed")
	}

	repositoryFactory, ok := repoFactories[repositoryName]

	if !ok {
		return nil, errors.New("Can't create Repository")
	}

	return repositoryFactory(conf)

	/*
	switch t {
	case "Mongo":
		return new(repository.MongoPage), nil
	default:
		return new(repository.FilesystemPage), nil
	}
	*/
}


// This implementation creates some sample files inside /tmp
// for the sole purpose to be used as a DB replacement.
// There is no proper error-checking in place nor specific cases
// are considered; a better implementation should
// consider adding proper validation, directory creation etc.

type FilesystemPage struct {
	*os.File
}

func InitFilesystem(path string, filename string) (*FilesystemPage, error) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}
	f, err := os.Create(path + filename)
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