package cacher

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func NewFileCache(path string) *FileCache {
	app := filepath.Dir(os.Args[0])
	cachepath := filepath.Join(app, path)
	ok, err := exists(cachepath)
	if err != nil {
	}
	if !ok {
		if err = os.Mkdir(cachepath, os.ModePerm); err != nil {
		}
	}

	return &FileCache{
		Path: cachepath,
	}
}

type FileCache struct {
	Path string
}

func (fc *FileCache) Get(key string) (*Item, error) {
	filename := filepath.Join(fc.Path, key)
	value, err := ioutil.ReadFile(filename)

	return &Item{
		Key:   key,
		Value: value,
	}, err
}

func (fc *FileCache) Set(key string, value []byte) (err error) {
	filename := filepath.Join(fc.Path, key)
	err = ioutil.WriteFile(filename, value, os.ModePerm)
	return err
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (fc *FileCache) Flush() (err error) {
	return os.RemoveAll(fc.Path)
}
